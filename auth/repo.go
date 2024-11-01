package auth

import (
	"database/sql"
	"go.uber.org/zap"
	"ndy/realworld-gin/logger"
)

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
	query := "SELECT id, email, password FROM users WHERE email = ?"
	err := repo.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		logger.Error("FindUserByEmail failed", zap.Error(err))
		return User{}, err
	}
	return user, nil
}

// FindProfileByUserID 는 주어진 사용자 ID에 해당하는 프로필을 반환합니다.
func (repo *MysqlRepo) FindProfileByUserID(userID int) (Profile, error) {
	var profile Profile
	query := "SELECT id, user_id, bio, image FROM profiles WHERE user_id = ?"
	err := repo.db.QueryRow(query, userID).Scan(&profile.Id, &profile.UserID, &profile.Bio, &profile.Image)
	if err != nil {
		logger.Error("FindProfileByUserID failed", zap.Error(err))
		return Profile{}, err
	}
	return profile, nil
}
