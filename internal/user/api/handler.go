package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/user/app"
	"ndy/realworld-gin/internal/user/dto"
	"ndy/realworld-gin/internal/util"
	"net/http"
	"strings"
)

func RegisterHandler(l *app.Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("rootData")
		var req dto.RegistrationRequest
		if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
			util.Log.Info("Error registering user", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
			return
		}

		if err := binding.Validator.ValidateStruct(req); err != nil {
			util.Log.Info("Error registering user", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		util.Log.Info("Registering user", zap.String("email", req.Email), zap.String("username", req.Username))

		_, err := (*l).Register(req.Username, req.Email, req.Password)
		if errors.Is(err, app.EmailAlreadyRegistered) {
			util.Log.Info("Email already registered", zap.String("email", req.Email))
			c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
			return
		}
		if err != nil {
			util.Log.Info("Error registering user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		util.Log.Info("Registered user", zap.String("email", req.Email), zap.String("username", req.Username))

		// Return the response
		c.Set("resp", dto.RegistrationResponse{
			Email:    req.Email,
			Username: req.Username,
			Token:    "",
			Bio:      "",
			Image:    "",
		})
	}
}

func GetCurrentUserHandler(l *app.Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the context
		userId, _ := c.Get("userId")
		profileId, _ := c.Get("profileId")
		util.Log.Info("Getting current user", zap.Int("userId", userId.(int)))

		// Get the current user
		resp, err := (*l).GetCurrentUser(userId.(int), profileId.(int))
		if err != nil {
			util.Log.Info("Error getting current user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		util.Log.Info("Got current user", zap.Int("userId", userId.(int)))

		// Return the response with the token
		resp.Token = strings.Replace(c.GetHeader("Authorization"), "Token ", "", 1)
		c.Set("resp", resp)
	}
}

func UpdateUserHandler(l *app.Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 데이터 획득
		data, _ := c.Get("rootData")
		var req dto.UpdateUserRequest
		if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
			util.Log.Info("Error updating user", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
			return
		}

		if err := binding.Validator.ValidateStruct(req); err != nil {
			util.Log.Info("Error updating user", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		util.Log.Info("Updating user", zap.String("email", req.Email), zap.String("username", req.Username))

		// 사용자 ID 획득
		userID, _ := c.Get("userId")
		profileId, _ := c.Get("profileId")
		ctx := context.WithValue(c, "userId", userID)
		ctx = context.WithValue(ctx, "profileId", profileId)

		// 사용자 정보 업데이트
		resp, err := (*l).UpdateUser(ctx, req.Email, req.Username, req.Password, req.Image, req.Bio)
		if err != nil {
			util.Log.Info("Error updating user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		util.Log.Info("Updated user", zap.String("email", req.Email), zap.String("username", req.Username))

		// Return the response with the token
		resp.Token = strings.Replace(c.GetHeader("Authorization"), "Token ", "", 1)
		c.Set("resp", resp)
	}
}
