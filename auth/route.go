package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
)

func Authentication(c *gin.Context) {
	userData, _ := c.Get("rootData")
	var req loginRequest
	if err := json.Unmarshal(userData.(json.RawMessage), &req); err != nil {
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
