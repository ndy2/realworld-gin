package table

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type UserRow struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserTable struct {
	DB *sqlx.DB
}

func NewUserTable(db *sqlx.DB) *UserTable {
	return &UserTable{DB: db}
}

func (t *UserTable) Insert(row UserRow) (int, error) {
	query := `INSERT INTO users (username, email, password, created_at, updated_at) 
			  VALUES (:username, :email, :password, :created_at, :updated_at)`
	result, err := t.DB.NamedExec(query, row)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (t *UserTable) FindByID(id int) (UserRow, error) {
	var row UserRow
	err := t.DB.Get(&row, "SELECT * FROM users WHERE id = ?", id)
	return row, err
}

func (t *UserTable) FindByEmail(email string) (UserRow, error) {
	var row UserRow
	err := t.DB.Get(&row, "SELECT * FROM users WHERE email = ?", email)
	return row, err
}

func (t *UserTable) Update(row UserRow) error {
	query := `UPDATE users SET username = :username, email = :email, password = :password, updated_at = :updated_at WHERE id = :id`
	_, err := t.DB.NamedExec(query, row)
	return err
}
