package app

import (
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/util"
	"time"
)

const jwtExpire = 72 * time.Hour
const jwtSecret = "your-256-bit-secret"

// Generate 함수는 사용자와 프로필 정보를 이용해 JWT 토큰을 생성합니다.
func Generate(userId, profileId int, username string) (string, error) {
	// 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"profileId": profileId,
		"username":  username,
		"exp":       time.Now().Add(jwtExpire).Unix(),
	})

	// 토큰을 서명합니다.
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		util.Log.Error("Failed to sign token", zap.Error(err))
		return "", err
	}

	// 생성된 토큰을 반환합니다.
	return tokenString, nil
}

func Verify(token string) (jwt.MapClaims, error) {
	// 토큰을 파싱합니다.
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		util.Log.Info("Failed to parse token", zap.Error(err))
		return nil, err
	}

	// 토큰이 유효한지 검증합니다.
	if !parsedToken.Valid {
		util.Log.Info("Invalid token")
		return nil, err
	}

	// 토큰에 포함된 클레임을 반환합니다.
	return parsedToken.Claims.(jwt.MapClaims), nil
}
