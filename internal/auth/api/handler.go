package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/auth/app"
	"ndy/realworld-gin/internal/auth/dto"
	"ndy/realworld-gin/internal/util"
	"net/http"
)

// AuthenticationHandler 는 사용자 인증을 처리합니다.
func AuthenticationHandler(l *app.Logic) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 데이터 획득
		data, _ := c.Get("rootData")
		var req dto.LoginRequest
		if err := json.Unmarshal(data.(json.RawMessage), &req); err != nil {
			util.Log.Info("Error authenticating user", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
			return
		}

		if err := binding.Validator.ValidateStruct(req); err != nil {
			util.Log.Info("Error authenticating user", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		util.Log.Info("Authenticating user", zap.String("email", req.Email))

		// 로그인, 토큰 생성
		resp, err := (*l).Login(req.Email, req.Password)

		// 예외 처리 및 응답 반환
		if errors.Is(err, app.ErrUserNotFound) || errors.Is(err, app.ErrPasswordMismatch) {
			util.Log.Info("Error authenticating user", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		if err != nil {
			util.Log.Error("Error authenticating user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		util.Log.Info("Authenticated user", zap.String("email", req.Email))

		// 응답 반환
		c.Set("resp", resp)
	}
}
