package auth

import "database/sql"

type MysqlRepo struct {
	db *sql.DB
}

// NewMysqlRepo 는 MysqlRepo 를 생성하고 반환합니다.
func NewMysqlRepo(db *sql.DB) *MysqlRepo {
	return &MysqlRepo{db: db}
}

// FindUserByEmail 는 주어진 사용자가 있는지 확인합니다.
func (repo *MysqlRepo) FindUserByEmail(email string) (User, error) {
	var user User
	query := "SELECT email, password FROM users WHERE email = ?"
	err := repo.db.QueryRow(query, email).Scan(&user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
