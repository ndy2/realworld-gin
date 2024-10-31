package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
)

// AuthenticationHandler 는 사용자 인증을 처리합니다.
func AuthenticationHandler(l *Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 데이터 획득
		data, _ := c.Get("rootData")
		var req LoginRequest
		if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
			return
		}
		log.Println("Authentication request:", req)

		// 로그인, 토큰 생성
		resp, err := l.Login(req.Email, req.Password)

		// 예외 처리 및 응답 반환
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrUserNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.Set("resp", resp)
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
