package app

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"ndy/realworld-gin/internal/auth/domain"
	"ndy/realworld-gin/internal/auth/dto"
	"testing"
)

func TestLogicImpl_Login(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		repo    domain.Repo
		args    args
		want    dto.LoginResponse
		wantErr bool
	}{
		{
			name: "valid login",
			repo: func() domain.Repo {
				mockRepo := domain.NewMockRepo(ctrl)
				mockRepo.EXPECT().FindUserByEmail("test@example.com").Return(domain.User{
					Id:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: string(hashedPassword), // bcrypt hash for "password"
				}, nil)
				mockRepo.EXPECT().FindProfileByUserID(1).Return(domain.Profile{
					Bio:   "This is a bio",
					Image: "http://example.com/image.jpg",
				}, nil)
				return mockRepo
			}(),
			args: args{
				email:    "test@example.com",
				password: "password",
			},
			want: dto.LoginResponse{
				Email:    "test@example.com",
				Token:    "generated_token", // Assuming the generate function returns "generated_token"
				Username: "testuser",
				Bio:      "This is a bio",
				Image:    "http://example.com/image.jpg",
			},
			wantErr: false,
		},
		{
			name: "user not found",
			repo: func() domain.Repo {
				mockRepo := domain.NewMockRepo(ctrl)
				mockRepo.EXPECT().FindUserByEmail("anonymous@example.com").Return(domain.User{}, ErrUserNotFound)
				return mockRepo
			}(),
			args: args{
				email:    "anonymous@example.com",
				password: "password",
			},
			want:    dto.LoginResponse{},
			wantErr: true,
		},
		{
			name: "password mismatch",
			repo: func() domain.Repo {
				mockRepo := domain.NewMockRepo(ctrl)
				mockRepo.EXPECT().FindUserByEmail("test@example.com").Return(domain.User{
					Id:       1,
					Email:    "test@example.com",
					Password: string(hashedPassword), // bcrypt hash for "password"
				}, nil)
				return mockRepo
			}(),
			args: args{
				email:    "test@example.com",
				password: "wrongpassword",
			},
			want:    dto.LoginResponse{},
			wantErr: true,
		},
		{
			name: "profile not found",
			repo: func() domain.Repo {
				mockRepo := domain.NewMockRepo(ctrl)
				mockRepo.EXPECT().FindUserByEmail("test@example.com").Return(domain.User{
					Id:       1,
					Email:    "test@example.com",
					Password: string(hashedPassword),
				}, nil)
				mockRepo.EXPECT().FindProfileByUserID(1).Return(domain.Profile{}, ErrProfileNotFound)
				return mockRepo
			}(),
			args: args{
				email:    "test@example.com",
				password: "password",
			},
			want:    dto.LoginResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LogicImpl{repo: tt.repo}
			got, err := l.Login(tt.args.email, tt.args.password)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.True(t, cmp.Equal(got, tt.want, cmpopts.IgnoreFields(dto.LoginResponse{}, "Token")))
		})
	}
}
