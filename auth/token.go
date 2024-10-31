package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

const jwtExpire = 72 * time.Hour
const jwtSecret = "your-256-bit-secret"

// generate 함수는 사용자와 프로필 정보를 이용해 JWT 토큰을 생성합니다.
func generate(u User, p Profile) (string, error) {
	// 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    u.Id,
		"profileId": p.Id,
		"username":  p.Username,
		"exp":       time.Now().Add(jwtExpire).Unix(),
	})

	// 토큰을 서명합니다.
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatalf("토큰 서명에 실패했습니다: %v", err)
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
		log.Printf("토큰 파싱에 실패했습니다: %v", err)
		return nil, err
	}

	// 토큰이 유효한지 검증합니다.
	if !parsedToken.Valid {
		log.Printf("유효하지 않은 토큰입니다.")
		return nil, err
	}

	// 토큰에 포함된 클레임을 반환합니다.
	return parsedToken.Claims.(jwt.MapClaims), nil
}
