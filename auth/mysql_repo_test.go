package auth

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMysqlRepo_FindUserByEmail(t *testing.T) {
	db, m, _ := NewMockDB()
	defer db.Close()

	// Mock a User row
	u1 := User{
		Id:       1,
		Email:    "user1@mail.com",
		Password: "password",
	}
	MockUserTable(m, u1.toRow())

	tests := []struct {
		name    string
		email   string
		want    User
		wantErr bool
	}{
		{
			name:    "user found 1",
			email:   "user1@mail.com",
			want:    u1,
			wantErr: false,
		},
		{
			name:    "user not found",
			email:   "no-user@mail.com",
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.FindUserByEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FindUserByEmail() got = %v, want %v", got, tt.want)
			}
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

	tests := []struct {
		name    string
		userId  int
		want    Profile
		wantErr bool
	}{
		{
			name:    "profile found 1",
			userId:  p1.UserID,
			want:    p1,
			wantErr: false,
		},
		{
			name:    "profile not found",
			userId:  3,
			want:    Profile{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MysqlRepo{db}
			got, err := repo.FindProfileByUserID(tt.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindProfileByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FindProfileByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (u User) toRow() UserRow {
	return UserRow{
		Id:       u.Id,
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
