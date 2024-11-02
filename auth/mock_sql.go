package auth

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
)

func NewMockDB() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		return nil, nil, err
	}
	return db, mock, nil
}

type UserRow struct {
	Id       int
	Email    string
	Password string
}

func MockUserTable(mock sqlmock.Sqlmock, users ...UserRow) {
	for _, user := range users {
		rows := sqlmock.NewRows([]string{"id", "email", "password"}).
			AddRow(user.Id, user.Email, user.Password)
		mock.ExpectQuery("SELECT id, email, password FROM users WHERE email = ?").
			WithArgs(user.Email).
			WillReturnRows(rows)
	}
}

type ProfileRow struct {
	Id     int
	UserID int
	Bio    string
	Image  string
}

func MockProfileTable(mock sqlmock.Sqlmock, profiles ...ProfileRow) {
	for _, profile := range profiles {
		rows := sqlmock.NewRows([]string{"id", "user_id", "bio", "image"}).
			AddRow(profile.Id, profile.UserID, profile.Bio, profile.Image)
		mock.ExpectQuery("SELECT id, user_id, bio, image FROM profiles WHERE user_id = ?").
			WithArgs(profile.UserID).
			WillReturnRows(rows)
	}
}
