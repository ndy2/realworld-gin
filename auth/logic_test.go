package auth

import (
	"go.uber.org/mock/gomock"
	"ndy/realworld-gin/logger"
	"reflect"
	"testing"
)

func init() {
	logger.InitLogger()
}

func XTestLogic_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//mockRepo := NewMockRepo(ctrl)

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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}
