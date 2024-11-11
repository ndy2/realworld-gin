package infra

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"ndy/realworld-gin/internal/auth/app"
	"ndy/realworld-gin/internal/auth/domain"
	"ndy/realworld-gin/internal/util/table"
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

func TestMysqlRepo_FindUserByEmail(t *testing.T) {
	// Mock a User row
	ur := table.UserRow{
		ID:       1,
		Username: "user1",
		Email:    "user1@mail.com",
		Password: "password",
	}
	MockUserTable(mock, ur)
	MockUserTableErrNoRow(mock, "no-user@mail.com")

	tests := []struct {
		name    string
		email   string
		want    domain.User
		wantErr error
	}{
		{
			name:    "user found 1",
			email:   "user1@mail.com",
			want:    toUser(ur),
			wantErr: nil,
		},
		{
			name:    "user not found",
			email:   "no-user@mail.com",
			want:    domain.User{},
			wantErr: app.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.FindUserByEmail(tt.email)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMysqlRepo_FindProfileByUserID(t *testing.T) {
	// Mock a Profile row
	pr := table.ProfileRow{
		ID:     1,
		UserID: 1,
		Bio:    "This is a bio",
		Image:  "http://example.com/image.jpg",
	}
	MockProfileTable(mock, pr)
	MockProfileTableErrNoRow(mock, 3)

	tests := []struct {
		name    string
		userId  int
		want    domain.Profile
		wantErr error
	}{
		{
			name:    "profile found 1",
			userId:  pr.UserID,
			want:    toProfile(pr),
			wantErr: nil,
		},
		{
			name:    "profile not found",
			userId:  3,
			want:    domain.Profile{},
			wantErr: app.ErrProfileNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.FindProfileByUserID(tt.userId)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
