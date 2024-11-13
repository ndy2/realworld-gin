package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"ndy/realworld-gin/internal/auth/app"
	"ndy/realworld-gin/internal/auth/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticationHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type want struct {
		status int
		resp   dto.LoginResponse
	}

	tests := []struct {
		name     string
		l        app.Logic
		rootData json.RawMessage
		want     want
	}{
		{
			name: "success",
			l: func() app.Logic {
				mockLogic := app.NewMockLogic(ctrl)
				mockLogic.EXPECT().Login("test@mail.com", "password").Return(dto.LoginResponse{
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
				resp: dto.LoginResponse{
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
			l:        app.NewMockLogic(ctrl),
			rootData: json.RawMessage(`{"email":"`),
			want: want{
				status: http.StatusBadRequest,
				resp:   dto.LoginResponse{}, // empty response
			},
		},
		{
			name: "password mismatch",
			l: func() app.Logic {
				mockLogic := app.NewMockLogic(ctrl)
				mockLogic.EXPECT().Login("test@mail.com", "wrongpassword").Return(dto.LoginResponse{},
					app.ErrPasswordMismatch)
				return mockLogic
			}(),
			rootData: json.RawMessage(`{"email":"test@mail.com","password":"wrongpassword"}`),
			want: want{
				status: http.StatusUnauthorized,
				resp:   dto.LoginResponse{}, // empty response
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
