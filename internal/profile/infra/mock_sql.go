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

type ProfileRow struct {
	Bio   string
	Image string
}

func MockProfile(mock sqlmock.Sqlmock, id int, row ProfileRow) {
	rows := sqlmock.NewRows([]string{"bio", "image"}).AddRow(row.Bio, row.Image)
	mock.ExpectQuery("SELECT bio, image FROM profiles WHERE id = ?").
		WithArgs(id).
		WillReturnRows(rows)
}
