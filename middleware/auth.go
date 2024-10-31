package middleware

import (
	"github.com/gin-gonic/gin"
	"ndy/realworld-gin/auth"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization 헤더로부터 토큰을 추출 (Bearer 토큰)
		token := strings.Replace(c.GetHeader("Authorization"), "Token ", "", 1)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 추출한 토큰을 검증
		claims, err := auth.Verify(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 검증된 토큰에서 사용자 ID, 프로필 ID 를 추출
		userID := int(claims["userId"].(float64))
		profileID := int(claims["profile"].(float64))

		// 추출한 사용자 ID, 프로필 ID 를 context 에 저장
		c.Set("userID", userID)
		c.Set("profileID", profileID)

		// 다음 핸들러로 요청을 전달
		c.Next()
	}
}