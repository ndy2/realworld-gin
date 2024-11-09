package infra

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/auth/app"
	"ndy/realworld-gin/internal/auth/domain"
	"ndy/realworld-gin/internal/util"
	"os"
)

type MysqlRepo struct {
	DB *sqlx.DB
}

// NewMysqlRepo 는 MysqlRepo 를 생성하고 반환합니다.
func NewMysqlRepo(dsn string) *MysqlRepo {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		util.Log.Fatal("NewMysqlRepo failed", zap.Error(err))
		os.Exit(1)
	}
	return &MysqlRepo{DB: db}
}

// FindUserByEmail 는 주어진 이메일로 사용자를 조회합니다.
func (repo *MysqlRepo) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, username, email, password FROM users WHERE email = ?"

	// SQLX의 Get을 사용하여 한 번에 데이터를 가져옵니다.
	err := repo.DB.Get(&user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 사용자 미발견시 에러 처리
			return domain.User{}, fmt.Errorf("%w: %v", app.ErrUserNotFound, err)
		}
		// 에러 로그 출력
		util.Log.Error("FindUserByEmail failed", zap.Error(err))
		return domain.User{}, err
	}
	return user, nil
}

// FindProfileByUserID 는 주어진 사용자 ID에 해당하는 프로필을 반환합니다.
func (repo *MysqlRepo) FindProfileByUserID(userID int) (domain.Profile, error) {
	var profile domain.Profile
	query := "SELECT id, user_id, bio, image FROM profiles WHERE user_id = ?"
	err := repo.DB.QueryRow(query, userID).Scan(&profile.Id, &profile.UserID, &profile.Bio, &profile.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Profile{}, fmt.Errorf("%w: %v", app.ErrProfileNotFound, err)
		}
		util.Log.Error("FindProfileByUserID failed", zap.Error(err))
		return domain.Profile{}, err
	}
	return profile, nil
}
