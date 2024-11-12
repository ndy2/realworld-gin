package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"ndy/realworld-gin/internal/profile/app"
	"ndy/realworld-gin/internal/profile/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProfileHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	gin.SetMode(gin.TestMode)

	type args struct {
		username      string
		authenticated bool
		userId        int
		profileId     int
	}

	type want struct {
		status int
		resp   dto.GetProfileResponse
	}

	tests := []struct {
		name string
		l    app.Logic
		args args
		want want
	}{
		{
			name: "authenticated user",
			l: func() app.Logic {
				mockLogic := app.NewMockLogic(ctrl)
				mockLogic.EXPECT().GetProfile(1, 1, "testuser", "targetuser").Return(dto.GetProfileResponse{
					Username:  "",
					Bio:       "",
					Image:     "",
					Following: false,
				}, nil)
				return mockLogic
			}(),
			args: args{
				username:      "targetuser",
				authenticated: true,
				userId:        1,
				profileId:     1,
			},
			want: want{
				status: 200,
				resp: dto.GetProfileResponse{
					Username:  "",
					Bio:       "",
					Image:     "",
					Following: false,
				},
			},
		},
		{
			name: "unauthenticated user",
			l: func() app.Logic {
				mockLogic := app.NewMockLogic(ctrl)
				mockLogic.EXPECT().GetProfile(0, 0, "", "targetuser").Return(dto.GetProfileResponse{
					Username:  "targetuser",
					Bio:       "targetbio",
					Image:     "",
					Following: false,
				}, nil)
				return mockLogic
			}(),
			args: args{
				username:      "targetuser",
				authenticated: false,
			},
			want: want{
				status: 200,
				resp: dto.GetProfileResponse{
					Username:  "targetuser",
					Bio:       "targetbio",
					Image:     "",
					Following: false,
				},
			},
		},
		{
			name: "profile not found",
			l: func() app.Logic {
				mockLogic := app.NewMockLogic(ctrl)
				mockLogic.EXPECT().GetProfile(0, 0, "", "targetuser").Return(dto.GetProfileResponse{}, app.ErrProfileNotFound)
				return mockLogic
			}(),
			args: args{
				username:      "targetuser",
				authenticated: false,
			},
			want: want{
				status: 404,
				resp:   dto.GetProfileResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.AddParam("username", "targetuser")
			c.Set("Authenticated", tt.args.authenticated)
			c.Set("userId", tt.args.userId)
			c.Set("profileId", tt.args.profileId)
			c.Set("username", "testuser")

			GetProfileHandler(&tt.l)(c)

			resp, _ := c.Get("resp")
			assert.Equal(t, tt.want.status, w.Code)
			if tt.want.status == http.StatusOK {
				assert.Equal(t, tt.want.resp, resp)
			}
		})
	}
}
