package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
)

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
