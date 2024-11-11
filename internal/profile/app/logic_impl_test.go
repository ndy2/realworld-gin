package app

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"ndy/realworld-gin/internal/profile/domain"
	"ndy/realworld-gin/internal/profile/dto"
	"testing"
)

func TestLogicImpl_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		currentUserId        int
		currentUserProfileId int
		currentUsername      string
		targetUsername       string
	}
	tests := []struct {
		name    string
		repo    domain.Repo
		args    args
		want    dto.GetProfileResponse
		wantErr bool
	}{
		{
			name: "Authenticated user is the same as the target user",
			repo: func() domain.Repo {
				mockRepo := domain.NewMockRepo(ctrl)
				mockRepo.EXPECT().FindProfile(1).Return(domain.Profile{
					Bio:   "This is a bio",
					Image: "http://example.com/image.jpg",
				}, nil)
				return mockRepo
			}(),
			args: args{
				currentUserId:        1,
				currentUserProfileId: 1,
				currentUsername:      "testuser",
				targetUsername:       "testuser",
			},
			want: dto.GetProfileResponse{
				Username:  "testuser",
				Bio:       "This is a bio",
				Image:     "http://example.com/image.jpg",
				Following: false,
			},
			wantErr: false,
		},
		{
			name: "Unauthenticated user is viewing the profile of another user",
			repo: func() domain.Repo {
				mockRepo := domain.NewMockRepo(ctrl)
				mockRepo.EXPECT().FindProfileByUsername("testuser").Return(domain.Profile{
					Bio:   "This is a bio",
					Image: "http://example.com/image.jpg",
				}, nil)
				return mockRepo
			}(),
			args: args{
				currentUserId:        0,
				currentUserProfileId: 0,
				currentUsername:      "",
				targetUsername:       "testuser",
			},
			want: dto.GetProfileResponse{
				Username:  "testuser",
				Bio:       "This is a bio",
				Image:     "http://example.com/image.jpg",
				Following: false,
			},
			wantErr: false,
		},
		{
			name: "Authenticated user is viewing the profile of another user",
			repo: func() domain.Repo {
				mockRepo := domain.NewMockRepo(ctrl)
				mockRepo.EXPECT().FindProfileWithFollowingByUsername("anotheruser", 1).Return(domain.Profile{
					Bio:   "This is a bio",
					Image: "http://example.com/image.jpg",
				}, domain.Following(true), nil)
				return mockRepo
			}(),
			args: args{
				currentUserId:        1,
				currentUserProfileId: 1,
				currentUsername:      "testuser",
				targetUsername:       "anotheruser",
			},
			want: dto.GetProfileResponse{
				Username:  "anotheruser",
				Bio:       "This is a bio",
				Image:     "http://example.com/image.jpg",
				Following: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LogicImpl{repo: tt.repo}
			got, err := l.GetProfile(tt.args.currentUserId, tt.args.currentUserProfileId, tt.args.currentUsername, tt.args.targetUsername)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
