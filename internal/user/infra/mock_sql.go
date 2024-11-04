package infra

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

func MockUserExistsByEmail(mock sqlmock.Sqlmock, email string, exists bool) {
	rows := sqlmock.NewRows([]string{"exists"}).AddRow(exists)
	mock.ExpectQuery("SELECT EXISTS \\(SELECT 1 FROM users WHERE email = \\?\\)").
		WithArgs(email).
		WillReturnRows(rows)
}
