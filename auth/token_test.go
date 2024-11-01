package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"ndy/realworld-gin/logger"
	"reflect"
	"strings"
	"testing"
	"time"
)

func init() {
	logger.InitLogger()
}

func TestVerify(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    jwt.MapClaims
		wantErr bool
	}{
		{
			name: "valid token",
			args: args{
				token: func() string {
					u := User{Id: 1, Email: "testuser@email.com"}
					p := Profile{Id: 1, Username: "testuser"}
					token, _ := generate(u, p)
					return token
				}(),
			},
			want: jwt.MapClaims{
				"userId":    float64(1),
				"profileId": float64(1),
				"username":  "testuser",
				"exp":       float64(time.Now().Add(jwtExpire).Unix()),
			},
			wantErr: false,
		},
		{
			name: "invalid token",
			args: args{
				token: "invalid.token.string",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Verify(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Verify() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generate(t *testing.T) {
	type args struct {
		u User
		p Profile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "generate valid token",
			args: args{
				u: User{Id: 1, Email: "testuser@email.com"},
				p: Profile{Id: 1, Username: "testuser"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generate(tt.args.u, tt.args.p)
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
			if claims["userId"] != float64(tt.args.u.Id) {
				t.Errorf("generate() userId = %v, want %v", claims["userId"], tt.args.u.Id)
			}
			if claims["profileId"] != float64(tt.args.p.Id) {
				t.Errorf("generate() profileId = %v, want %v", claims["profileId"], tt.args.p.Id)
			}
			if claims["username"] != tt.args.p.Username {
				t.Errorf("generate() username = %v, want %v", claims["username"], tt.args.p.Username)
			}
		})
	}
}
