package app

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"ndy/realworld-gin/internal/user/domain"
	"testing"
)

func TestLogicImpl_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := domain.NewMockRepo(ctrl)

	type args struct {
		username string
		email    string
		password string
	}
	tests := []struct {
		name    string
		repo    domain.Repo
		args    args
		want    int
		wantErr error
	}{
		{
			name: "valid register",
			repo: func() domain.Repo {
				mockRepo.EXPECT().CheckUserExists("jake@jake.jake").Return(false, nil).AnyTimes()
				mockRepo.EXPECT().InsertUser(userMatcher{
					Username: "jake",
					Email:    "jake@jake.jake",
					Password: "password",
				}).Return(1, nil).AnyTimes()
				return mockRepo
			}(),
			args: args{
				username: "jake",
				email:    "jake@jake.jake",
				password: "password",
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "email already registered",
			repo: func() domain.Repo {
				mockRepo.EXPECT().CheckUserExists("jake2@jake.jake").Return(true, nil).AnyTimes()
				return mockRepo
			}(),
			args: args{
				username: "jake",
				email:    "jake2@jake.jake",
				password: "password",
			},
			want:    0,
			wantErr: EmailAlreadyRegistered,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LogicImpl{repo: tt.repo}
			got, err := l.Register(tt.args.username, tt.args.email, tt.args.password)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

type userMatcher struct {
	Username string
	Email    string
	Password string
}

func (um userMatcher) Matches(x interface{}) bool {
	u, _ := x.(domain.User)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(um.Password))
	return err == nil && u.Username == um.Username && u.Email == um.Email
}

func (um userMatcher) String() string {
	return "is user with bcrypt hashed password"
}

func TestLogicImpl_GetCurrentUser(t *testing.T) {
	// TODO: Add test cases.
}

func TestLogicImpl_UpdateUser(t *testing.T) {
	// TODO: Add test cases.
}
