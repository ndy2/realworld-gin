package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"ndy/realworld-gin/logger"
	"net/http"
	"strings"
)

func RegisterHandler(l *Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("rootData")
		var req RegistrationRequest
		if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
			logger.Info("Error registering user", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
			return
		}
		logger.Info("Registering user", zap.String("email", req.Email), zap.String("username", req.Username))

		_, err := l.Register(req.Username, req.Email, req.Password)
		if err != nil {
			logger.Info("Error registering user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logger.Info("Registered user", zap.String("email", req.Email), zap.String("username", req.Username))

		// Return the response
		c.Set("resp", RegistrationResponse{
			Email:    req.Email,
			Username: req.Username,
			Token:    "",
			Bio:      "",
			Image:    "",
		})
	}
}

type RegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegistrationResponse userResponse

func GetCurrentUserHandler(l *Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the context
		userId, _ := c.Get("userId")
		profileId, _ := c.Get("profileId")
		logger.Info("Getting current user", zap.Int("userId", userId.(int)))

		// Get the current user
		resp, err := l.GetCurrentUser(userId.(int), profileId.(int))
		if err != nil {
			logger.Info("Error getting current user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logger.Info("Got current user", zap.Int("userId", userId.(int)))

		// Return the response with the token
		resp.Token = strings.Replace(c.GetHeader("Authorization"), "Token ", "", 1)
		c.Set("resp", resp)
	}
}

type GetCurrentUserResponse userResponse

func UpdateUserHandler(l *Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 데이터 획득
		data, _ := c.Get("rootData")
		var req UpdateUserRequest
		if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
			logger.Info("Error updating user", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
			return
		}
		logger.Info("Updating user", zap.String("email", req.Email), zap.String("username", req.Username))

		// 사용자 ID 획득
		userID, _ := c.Get("userId")
		profileId, _ := c.Get("profileId")
		ctx := context.WithValue(c, "userId", userID)
		ctx = context.WithValue(ctx, "profileId", profileId)

		// 사용자 정보 업데이트
		resp, err := l.UpdateUser(ctx, req.Email, req.Username, req.Password, req.Image, req.Bio)
		if err != nil {
			logger.Info("Error updating user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logger.Info("Updated user", zap.String("email", req.Email), zap.String("username", req.Username))

		// Return the response with the token
		resp.Token = strings.Replace(c.GetHeader("Authorization"), "Token ", "", 1)
		c.Set("resp", resp)
	}
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Image    string `json:"image"`
	Bio      string `json:"bio"`
}
type UpdateUserResponse userResponse

type userResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
