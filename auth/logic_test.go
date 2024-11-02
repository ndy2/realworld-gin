package auth

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"ndy/realworld-gin/logger"
	"testing"
)

func init() {
	logger.InitLogger()
}

func TestLogic_Login(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo Repo
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    LoginResponse
		wantErr bool
	}{
		{
			name: "valid login",
			fields: fields{
				repo: func() Repo {
					mockRepo := NewMockRepo(ctrl)
					mockRepo.EXPECT().FindUserByEmail("test@example.com").Return(User{
						Id:       1,
						Email:    "test@example.com",
						Password: string(hashedPassword), // bcrypt hash for "password"
					}, nil)
					mockRepo.EXPECT().FindProfileByUserID(1).Return(Profile{
						Username: "testuser",
						Bio:      "This is a bio",
						Image:    "http://example.com/image.jpg",
					}, nil)
					return mockRepo
				}(),
			},
			args: args{
				email:    "test@example.com",
				password: "password",
			},
			want: LoginResponse{
				Email:    "test@example.com",
				Token:    "generated_token", // Assuming the generate function returns "generated_token"
				Username: "testuser",
				Bio:      "This is a bio",
				Image:    "http://example.com/image.jpg",
			},
			wantErr: false,
		},
		{
			name: "anonymous email",
			fields: fields{
				repo: func() Repo {
					mockRepo := NewMockRepo(ctrl)
					mockRepo.EXPECT().FindUserByEmail("anonymous@example.com").Return(User{}, errors.New("user not found"))
					return mockRepo
				}(),
			},
			args: args{
				email:    "anonymous@example.com",
				password: "password",
			},
			want:    LoginResponse{},
			wantErr: true,
		},
		{
			name: "password mismatch",
			fields: fields{
				repo: func() Repo {
					mockRepo := NewMockRepo(ctrl)
					mockRepo.EXPECT().FindUserByEmail("test@example.com").Return(User{
						Id:       1,
						Email:    "test@example.com",
						Password: string(hashedPassword), // bcrypt hash for "password"
					}, nil)
					return mockRepo
				}(),
			},
			args: args{
				email:    "test@example.com",
				password: "wrongpassword",
			},
			want:    LoginResponse{},
			wantErr: true,
		},
		{
			name: "profile not found",
			fields: fields{
				repo: func() Repo {
					mockRepo := NewMockRepo(ctrl)
					mockRepo.EXPECT().FindUserByEmail("test@example.com").Return(User{
						Id:       1,
						Email:    "test@example.com",
						Password: string(hashedPassword),
					}, nil)
					mockRepo.EXPECT().FindProfileByUserID(1).Return(Profile{}, errors.New("profile not found"))
					return mockRepo
				}(),
			},
			args: args{
				email:    "test@example.com",
				password: "password",
			},
			want:    LoginResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logic{
				repo: tt.fields.repo,
			}
			got, err := l.Login(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, cmpopts.IgnoreFields(LoginResponse{}, "Token")) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}
