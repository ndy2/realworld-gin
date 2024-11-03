package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticationHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	gin.SetMode(gin.TestMode)
	type want struct {
		status int
		resp   LoginResponse
	}

	tests := []struct {
		name     string
		l        Logic
		rootData json.RawMessage
		want     want
	}{
		{
			name: "success",
			l: func() Logic {
				mockLogic := NewMockLogic(ctrl)
				mockLogic.EXPECT().Login("test@mail.com", "password").Return(LoginResponse{
					Email:    "test@mail.com",
					Token:    "generated_token",
					Username: "testuser",
					Bio:      "This is a bio",
					Image:    "http://example.com/image.jpg",
				}, nil)
				return mockLogic
			}(),
			rootData: json.RawMessage(`{"email":"test@mail.com","password":"password"}`),
			want: want{
				status: http.StatusOK,
				resp: LoginResponse{
					Email:    "test@mail.com",
					Token:    "generated_token",
					Username: "testuser",
					Bio:      "This is a bio",
					Image:    "http://example.com/image.jpg",
				},
			},
		},
		{
			name:     "invalid root data",
			l:        NewMockLogic(ctrl),
			rootData: json.RawMessage(`{"email":"`),
			want: want{
				status: http.StatusBadRequest,
				resp:   LoginResponse{}, // empty response
			},
		},
		{
			name: "password mismatch",
			l: func() Logic {
				mockLogic := NewMockLogic(ctrl)
				mockLogic.EXPECT().Login("test@mail.com", "wrongpassword").Return(LoginResponse{},
					ErrPasswordMismatch)
				return mockLogic
			}(),
			rootData: json.RawMessage(`{"email":"test@mail.com","password":"wrongpassword"}`),
			want: want{
				status: http.StatusUnauthorized,
				resp:   LoginResponse{}, // empty response
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Set("rootData", tt.rootData)

			AuthenticationHandler(&tt.l)(c)

			resp, _ := c.Get("resp")
			assert.Equal(t, tt.want.status, w.Code)
			if tt.want.status == http.StatusOK {
				assert.Equal(t, tt.want.resp, resp)
			}
		})
	}
}
