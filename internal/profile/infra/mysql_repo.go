package infra

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/profile/domain"
	"ndy/realworld-gin/internal/util"
	"os"
)

type MysqlRepo struct {
	db *sqlx.DB
}

// NewMysqlRepo 는 MysqlRepo 를 생성하고 반환합니다.
func NewMysqlRepo(dsn string) *MysqlRepo {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		util.Log.Fatal("NewMysqlRepo failed", zap.Error(err))
		os.Exit(1)
	}
	return &MysqlRepo{db: db}
}

// FindProfile 는 주어진 사용자 ID에 해당하는 프로필을 반환합니다.
func (repo *MysqlRepo) FindProfile(profileID int) (domain.Profile, error) {
	var profile domain.Profile
	query := "SELECT bio, image FROM profiles WHERE id = ?"
	err := repo.db.QueryRow(query, profileID).Scan(&profile.Bio, &profile.Image)
	if err != nil {
		util.Log.Error("FindProfileByID failed", zap.Error(err))
		return domain.Profile{}, err
	}
	return profile, nil
}

// FindProfileByUsername 는 주어진 사용자 이름에 해당하는 프로필을 반환합니다.
func (repo *MysqlRepo) FindProfileByUsername(username string) (domain.Profile, error) {
	var profile domain.Profile
	query := "SELECT bio, image FROM profiles WHERE username = ?"
	err := repo.db.QueryRow(query, username).Scan(&profile.Bio, &profile.Image)
	if err != nil {
		util.Log.Error("FindProfileByUsername failed", zap.Error(err))
		return domain.Profile{}, err
	}
	return profile, nil
}

// FindProfileWithFollowingByUsername 는 주어진 사용자 이름에 해당하는 프로필과 팔로잉 여부를 반환합니다.
func (repo *MysqlRepo) FindProfileWithFollowingByUsername(username string, currentUserId int) (domain.Profile, domain.Following, error) {
	var targetUserId int
	var profile domain.Profile
	var following bool
	query := "SELECT userId, bio, image FROM profiles WHERE username = ?"
	err := repo.db.QueryRow(query, username).Scan(&targetUserId, &profile.Bio, &profile.Image)
	if err != nil {
		util.Log.Error("FindProfileByUsername failed", zap.Error(err))
		return domain.Profile{}, false, err
	}

	query = "SELECT EXISTS (SELECT 1 FROM followers WHERE user_id = ? AND follower_id = ?)"
	err = repo.db.QueryRow(query, currentUserId, targetUserId).Scan(&following)
	if err != nil {
		util.Log.Error("FindProfileWithFollowingByUsername failed", zap.Error(err))
		return domain.Profile{}, false, err
	}

	return profile, domain.Following(following), nil
}
