package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/go-cmp/cmp"
	"ndy/realworld-gin/logger"
	"strings"
	"testing"
	"time"
)

func init() {
	logger.InitLogger()
}

func TestVerify(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		want    jwt.MapClaims
		wantErr bool
	}{
		{
			name: "valid token",
			token: func() string {
				u := User{Id: 1, Email: "testuser@email.com"}
				p := Profile{Id: 1, Username: "testuser"}
				token, _ := generate(u, p)
				return token
			}(),
			want: jwt.MapClaims{
				"userId":    float64(1),
				"profileId": float64(1),
				"username":  "testuser",
				"exp":       float64(time.Now().Add(jwtExpire).Unix()),
			},
			wantErr: false,
		},
		{
			name:    "invalid token",
			token:   "invalid.token.string",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Verify(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Verify() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generate(t *testing.T) {
	tests := []struct {
		name    string
		u       User
		p       Profile
		wantErr bool
	}{
		{
			name:    "generate valid token",
			u:       User{Id: 1, Email: "testuser@email.com"},
			p:       Profile{Id: 1, Username: "testuser"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generate(tt.u, tt.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.Count(got, ".") != 2 {
				t.Errorf("generate() failed to generate token")
				return
			}

			claims, err := Verify(got)
			if err != nil {
				t.Errorf("generate() failed to verify token = %v", err)
				return
			}
			if claims["userId"] != float64(tt.u.Id) {
				t.Errorf("generate() userId = %v, want %v", claims["userId"], tt.u.Id)
			}
			if claims["profileId"] != float64(tt.p.Id) {
				t.Errorf("generate() profileId = %v, want %v", claims["profileId"], tt.p.Id)
			}
			if claims["username"] != tt.p.Username {
				t.Errorf("generate() username = %v, want %v", claims["username"], tt.p.Username)
			}
		})
	}
}
