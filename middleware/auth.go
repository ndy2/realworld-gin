package middleware

import (
	"github.com/gin-gonic/gin"
	"ndy/realworld-gin/auth"
	"net/http"
	"strings"
)

// Auth 는 Authorization 헤더로부터 토큰을 추출하고, 토큰을 검증하는 미들웨어를 생성합니다.
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
		userId := int(claims["userId"].(float64))
		username := claims["username"].(string)
		profileId := int(claims["profileId"].(float64))

		// 추출한 사용자 ID, 프로필 ID 를 context 에 저장
		c.Set("userId", userId)
		c.Set("username", username)
		c.Set("profileId", profileId)

		// 다음 핸들러로 요청을 전달
		c.Next()
	}
}
