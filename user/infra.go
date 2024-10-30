package user

import "database/sql"

type MysqlUserRepo struct {
	db *sql.DB
}

// NewMysqlUserRepo 는 MysqlUserRepo 를 생성하고 반환합니다.
func NewMysqlUserRepo(db *sql.DB) *MysqlUserRepo {
	return &MysqlUserRepo{db: db}
}

// CheckUserExists 는 주어진 이메일을 가진 사용자가 있는지 확인합니다.
func (repo *MysqlUserRepo) CheckUserExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)"
	err := repo.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// InsertUser 는 새로운 사용자를 데이터베이스에 등록하고 새 사용자 ID를 반환합니다.
func (repo *MysqlUserRepo) InsertUser(u User) (int, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := repo.db.Exec(query, u.Username, u.Email, u.Password)
	if err != nil {
		return 0, err
	}

	// 새로 삽입된 사용자 ID를 가져옵니다.
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(userID), nil
}
