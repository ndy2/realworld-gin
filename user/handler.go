package user

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"strings"
)

func RegisterHandler(l *Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("rootData")
		var req RegistrationRequest
		if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
			return
		}
		log.Println("Registration request:", req)

		_, err := l.Register(req.Username, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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
		userID, _ := c.Get("userID")
		profileID, _ := c.Get("profileID")

		resp, err := l.GetCurrentUser(userID.(int), profileID.(int))
		resp.Token = strings.Replace(c.GetHeader("Authorization"), "Token ", "", 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Set("resp", resp)
	}
}

type GetCurrentUserResponse userResponse

func UpdateUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		type updateUserRequest struct {
			Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
			Image    string `json:"image"`
			Bio      string `json:"bio"`
		}
		type updateUserResponse userResponse
	}
}

type userResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
