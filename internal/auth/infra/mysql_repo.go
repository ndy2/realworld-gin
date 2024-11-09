package infra

import (
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/auth/app"
	"ndy/realworld-gin/internal/auth/domain"
	"ndy/realworld-gin/internal/util"
)

type MysqlRepo struct {
	db *sql.DB
}

// NewMysqlRepo 는 MysqlRepo 를 생성하고 반환합니다.
func NewMysqlRepo(db *sql.DB) *MysqlRepo {
	return &MysqlRepo{db: db}
}

// FindUserByEmail 는 주어진 사용자가 있는지 확인합니다.
func (repo *MysqlRepo) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, username, email, password FROM users WHERE email = ?"
	err := repo.db.QueryRow(query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, fmt.Errorf("%w: %v", app.ErrUserNotFound, err)
		}
		util.Log.Error("FindUserByEmail failed", zap.Error(err))
		return domain.User{}, err
	}
	return user, nil
}

// FindProfileByUserID 는 주어진 사용자 ID에 해당하는 프로필을 반환합니다.
func (repo *MysqlRepo) FindProfileByUserID(userID int) (domain.Profile, error) {
	var profile domain.Profile
	query := "SELECT id, user_id, bio, image FROM profiles WHERE user_id = ?"
	err := repo.db.QueryRow(query, userID).Scan(&profile.Id, &profile.UserID, &profile.Bio, &profile.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Profile{}, fmt.Errorf("%w: %v", app.ErrProfileNotFound, err)
		}
		util.Log.Error("FindProfileByUserID failed", zap.Error(err))
		return domain.Profile{}, err
	}
	return profile, nil
}
