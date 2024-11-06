package infra

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"ndy/realworld-gin/internal/profile/domain"
	"ndy/realworld-gin/internal/util"
	"os"
	"testing"
)

var db *sql.DB
var mock sqlmock.Sqlmock

func init() {
	util.InitLogger()
}

func TestMain(m *testing.M) {
	db, mock, _ = NewMockDB()
	defer db.Close()

	code := m.Run()

	os.Exit(code)
}

func TestMysqlRepo_FindProfile(t *testing.T) {
	// Mock a Profile Exists Query
	MockProfile(mock, 1, ProfileRow{
		Bio:   "bio1",
		Image: "image1",
	})

	tests := []struct {
		name      string
		profileID int
		want      domain.Profile
		wantErr   bool
	}{
		{
			name:      "profile exists",
			profileID: 1,
			want: domain.Profile{
				Bio:   "bio1",
				Image: "image1",
			},
			wantErr: false,
		},
		{
			name:      "profile not exists",
			profileID: 2,
			want:      domain.Profile{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.FindProfile(tt.profileID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMysqlRepo_FindProfileWithFollowingByUsername(t *testing.T) {
	MockProfileWithUserId(mock, "user2", 2, ProfileRow{
		Bio:   "bio2",
		Image: "image2",
	})
	MockFollowing(mock, 1, 2, true)

	type args struct {
		username      string
		currentUserId int
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Profile
		want1   domain.Following
		wantErr bool
	}{
		{
			name: "user following",
			args: args{
				username:      "user2",
				currentUserId: 1,
			},
			want: domain.Profile{
				Bio:   "bio2",
				Image: "image2",
			}, // has userId 2
			want1:   true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, got1, err := repo.FindProfileWithFollowingByUsername(tt.args.username, tt.args.currentUserId)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
