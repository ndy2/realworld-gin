package user

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
)

func RegisterHandler(l *Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		type registrationRequest struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		type registrationResponse userResponse

		data, _ := c.Get("rootData")
		var req registrationRequest
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
		c.Set("resp", registrationResponse{
			Email:    req.Email,
			Username: req.Username,
			Token:    "some.jwt.token",
			Bio:      "",
			Image:    "",
		})
	}
}
