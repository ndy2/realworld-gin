package infra

import (
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	auth3 "ndy/realworld-gin/internal/auth/app"
	auth2 "ndy/realworld-gin/internal/auth/domain"
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
func (repo *MysqlRepo) FindUserByEmail(email string) (auth2.User, error) {
	var user auth2.User
	query := "SELECT id, username, email, password FROM users WHERE email = ?"
	err := repo.db.QueryRow(query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return auth2.User{}, fmt.Errorf("%w: %v", auth3.ErrUserNotFound, err)
		}
		util.Log.Error("FindUserByEmail failed", zap.Error(err))
		return auth2.User{}, err
	}
	return user, nil
}

// FindProfileByUserID 는 주어진 사용자 ID에 해당하는 프로필을 반환합니다.
func (repo *MysqlRepo) FindProfileByUserID(userID int) (auth2.Profile, error) {
	var profile auth2.Profile
	query := "SELECT id, user_id, bio, image FROM profiles WHERE user_id = ?"
	err := repo.db.QueryRow(query, userID).Scan(&profile.Id, &profile.UserID, &profile.Bio, &profile.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return auth2.Profile{}, fmt.Errorf("%w: %v", auth3.ErrProfileNotFound, err)
		}
		util.Log.Error("FindProfileByUserID failed", zap.Error(err))
		return auth2.Profile{}, err
	}
	return profile, nil
}
