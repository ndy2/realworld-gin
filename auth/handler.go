package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"ndy/realworld-gin/logger"
	"net/http"
)

// AuthenticationHandler 는 사용자 인증을 처리합니다.
func AuthenticationHandler(l *Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 데이터 획득
		data, _ := c.Get("rootData")
		var req LoginRequest
		if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
			logger.Log.Info("Error authenticating user", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
			return
		}
		logger.Log.Info("Authenticating user", zap.String("email", req.Email))

		// 로그인, 토큰 생성
		resp, err := (*l).Login(req.Email, req.Password)

		// 예외 처리 및 응답 반환
		if errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrPasswordMismatch) {
			logger.Log.Info("Error authenticating user", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		if err != nil {
			logger.Log.Error("Error authenticating user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logger.Log.Info("Authenticated user", zap.String("email", req.Email))

		// 응답 반환
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
