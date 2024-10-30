package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
)

func AuthenticationHandler(l *Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		type loginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		type loginResponse struct {
			Email    string `json:"email"`
			Token    string `json:"token"`
			Username string `json:"username"`
			Bio      string `json:"bio"`
			Image    string `json:"image"`
		}

		data, _ := c.Get("rootData")
		var req loginRequest
		if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
			return
		}
		log.Println("Authentication request:", req)

		token, err := l.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrUserNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		resp := loginResponse{
			Email:    req.Email,
			Token:    token,
			Username: "some username",
			Bio:      "some bio",
			Image:    "some image url",
		}
		c.Set("resp", resp)
	}
}

func Authentication(c *gin.Context) {
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type loginResponse struct {
		Email    string `json:"email"`
		Token    string `json:"token"`
		Username string `json:"username"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	}

	data, _ := c.Get("rootData")
	var req loginRequest
	if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
		return
	}
	log.Println("Authentication request:", req)

	resp := loginResponse{
		Email:    req.Email,
		Token:    "some.jwt.token",
		Username: "some username",
		Bio:      "some bio",
		Image:    "some image url",
	}
	c.Set("resp", resp)
}
