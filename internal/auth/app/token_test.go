package app

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
	"time"
)

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
				token, _ := Generate(1, 1, "testuser")
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

func Test_Generate(t *testing.T) {
	tests := []struct {
		name      string
		userId    int
		profileId int
		username  string
		wantErr   bool
	}{
		{
			name:      "Generate valid token",
			userId:    1,
			profileId: 1,
			username:  "testuser",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.userId, tt.profileId, tt.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.Count(got, ".") != 2 {
				t.Errorf("Generate() failed to Generate token")
				return
			}

			claims, err := Verify(got)
			if err != nil {
				t.Errorf("Generate() failed to verify token = %v", err)
				return
			}
			if claims["userId"] != float64(tt.userId) {
				t.Errorf("Generate() userId = %v, want %v", claims["userId"], tt.userId)
			}
			if claims["profileId"] != float64(tt.profileId) {
				t.Errorf("Generate() profileId = %v, want %v", claims["profileId"], tt.profileId)
			}
			if claims["username"] != tt.username {
				t.Errorf("Generate() username = %v, want %v", claims["username"], tt.username)
			}
		})
	}
}
