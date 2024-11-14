package infra

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"ndy/realworld-gin/internal/user/domain"
	"os"
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
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMysqlRepo_InsertUserProfile(t *testing.T) {
	mock.ExpectExec("INSERT INTO users").
		WithArgs("test", "test@mail.com", "password", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1)) // (ID, RowsAffected)

	mock.ExpectExec("INSERT INTO profiles").
		WithArgs(1, "", "", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1)) // (ID, RowsAffected)

	type args struct {
		up domain.UserProfile
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "insert user",
			args: args{up: domain.UserProfile{
				User: domain.User{
					Username: "test",
					Email:    "test@mail.com",
					Password: "password",
				},
			}},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.InsertUserProfile(tt.args.up)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
		err := mock.ExpectationsWereMet()
		assert.NoError(t, err)
	}
}

func TestMysqlRepo_FindUserByID(t *testing.T) {
	mock.ExpectQuery("SELECT username, email FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"username", "email"}).AddRow("test", "test@mail.com"))

	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		args    args
		want    domain.User
		wantErr bool
	}{
		{
			name:    "find user",
			args:    args{userID: 1},
			want:    domain.User{Username: "test", Email: "test@mail.com"},
			wantErr: false,
		},
		{
			name:    "user not found",
			args:    args{userID: 2},
			want:    domain.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.FindUserByID(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMysqlRepo_UpdateProfile(t *testing.T) {
	mock.ExpectExec("UPDATE profiles").
		WithArgs("Updated bio", "Updated bio", "updated_image.jpg", "updated_image.jpg", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE profiles").
		WithArgs("Updated bio", "Updated bio", "", "", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	type args struct {
		profileId int
		profile   domain.Profile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update profile",
			args: args{profileId: 1, profile: domain.Profile{
				Bio:   "Updated bio",
				Image: "updated_image.jpg",
			}},
			wantErr: false,
		},
		{
			name: "update profile bio only",
			args: args{profileId: 1, profile: domain.Profile{
				Bio:   "Updated bio",
				Image: "",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			err := repo.UpdateProfile(tt.args.profileId, tt.args.profile)
			assert.Equal(t, tt.wantErr, err != nil)
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
			err := repo.UpdateUser(tt.args.userId, tt.args.user)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
