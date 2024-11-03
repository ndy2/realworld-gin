package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMysqlRepo_FindUserByEmail(t *testing.T) {
	db, m, _ := NewMockDB()
	defer db.Close()

	// Mock a User row
	u1 := User{
		Id:       1,
		Username: "user1",
		Email:    "user1@mail.com",
		Password: "password",
	}
	MockUserTable(m, u1.toRow())
	MockUserTableErrNoRow(m, "no-user@mail.com")

	tests := []struct {
		name    string
		email   string
		want    User
		wantErr error
	}{
		{
			name:    "user found 1",
			email:   "user1@mail.com",
			want:    u1,
			wantErr: nil,
		},
		{
			name:    "user not found",
			email:   "no-user@mail.com",
			want:    User{},
			wantErr: ErrUserNotFound,
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
	db, m, _ := NewMockDB()
	defer db.Close()

	// Mock a Profile row
	p1 := Profile{
		Id:     1,
		UserID: 1,
		Bio:    "This is a bio",
		Image:  "http://example.com/image.jpg",
	}
	MockProfileTable(m, p1.toRow())
	MockProfileTableErrNoRow(m, 3)

	tests := []struct {
		name    string
		userId  int
		want    Profile
		wantErr error
	}{
		{
			name:    "profile found 1",
			userId:  p1.UserID,
			want:    p1,
			wantErr: nil,
		},
		{
			name:    "profile not found",
			userId:  3,
			want:    Profile{},
			wantErr: ErrProfileNotFound,
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

func (u User) toRow() UserRow {
	return UserRow{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (p Profile) toRow() ProfileRow {
	return ProfileRow{
		Id:     p.Id,
		UserID: p.UserID,
		Bio:    p.Bio,
		Image:  p.Image,
	}
}
