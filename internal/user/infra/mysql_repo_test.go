package infra

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"ndy/realworld-gin/internal/user/domain"
	"os"
	"reflect"
	"testing"
)

var db *sqlx.DB
var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, mock, _ = NewMockDB()
	defer db.Close()

	code := m.Run()

	os.Exit(code)
}

func TestMysqlRepo_CheckUserExists(t *testing.T) {
	// Mock a User Exists Query
	MockUserExistsByEmail(mock, "user1@mail.com", true)
	MockUserExistsByEmail(mock, "not-exists@mail.com", false)

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "user exists",
			args:    args{email: "user1@mail.com"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "user not exists",
			args:    args{email: "not-exists@mail.com"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "some DB error",
			args:    args{email: "some DB error"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.CheckUserExists(tt.args.email)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMysqlRepo_FindProfileByID(t *testing.T) {
	type args struct {
		profileID int
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Profile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.FindProfileByID(tt.args.profileID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindProfileByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindProfileByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMysqlRepo_FindUserByID(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.FindUserByID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMysqlRepo_InsertUser(t *testing.T) {
	mock.ExpectExec("INSERT INTO users").
		WithArgs("test", "test@mail.com", "password", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1)) // (ID, RowsAffected)

	mock.ExpectExec("INSERT INTO profiles").
		WithArgs(1, "", "", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1)) // (ID, RowsAffected)

	type args struct {
		u domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "insert user",
			args: args{u: domain.User{
				Username: "test",
				Email:    "test@mail.com",
				Password: "password",
			}},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.InsertUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertUser() got = %v, want %v", got, tt.want)
			}
		})
		err := mock.ExpectationsWereMet()
		assert.NoError(t, err)
	}
}

func TestMysqlRepo_UpdateProfile(t *testing.T) {
	type args struct {
		profileId int
		profile   domain.Profile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			if err := repo.UpdateProfile(tt.args.profileId, tt.args.profile); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMysqlRepo_UpdateUser(t *testing.T) {
	type args struct {
		userId int
		user   domain.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			if err := repo.UpdateUser(tt.args.userId, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
