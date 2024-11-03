package profile

import (
	"database/sql"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/util"
)

type MysqlRepo struct {
	db *sql.DB
}

// NewMysqlRepo 는 MysqlRepo 를 생성하고 반환합니다.
func NewMysqlRepo(db *sql.DB) *MysqlRepo {
	return &MysqlRepo{db: db}
}

// FindProfile 는 주어진 사용자 ID에 해당하는 프로필을 반환합니다.
func (repo *MysqlRepo) FindProfile(profileID int) (Profile, error) {
	var profile Profile
	query := "SELECT bio, image FROM profiles WHERE id = ?"
	err := repo.db.QueryRow(query, profileID).Scan(&profile.Bio, &profile.Image)
	if err != nil {
		util.Log.Error("FindProfileByID failed", zap.Error(err))
		return Profile{}, err
	}
	return profile, nil
}
