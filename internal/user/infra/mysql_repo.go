package infra

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/user/domain"
	"ndy/realworld-gin/internal/util"
	"ndy/realworld-gin/internal/util/table"
	"os"
	"time"
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

// CheckUserExists 는 주어진 이메일을 가진 사용자가 있는지 확인합니다.
func (repo *MysqlRepo) CheckUserExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)"
	err := repo.DB.Get(&exists, query, email)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		util.Log.Error("CheckUserExists failed", zap.Error(err))
		return false, err
	}
	return exists, nil
}

// InsertUser 는 새로운 사용자를 데이터베이스에 등록하고 새 사용자 ID를 반환합니다.
func (repo *MysqlRepo) InsertUser(u domain.User) (int, error) {
	query := `INSERT INTO users (username, email, password, created_at, updated_at) 
			  VALUES (:username, :email, :password, :created_at, :updated_at)`
	result, err := repo.DB.NamedExec(query, table.UserRow{
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		util.Log.Error("InsertUser failed", zap.Error(err))
		return 0, err
	}

	// 새로 삽입된 사용자 ID를 가져옵니다.
	userID, err := result.LastInsertId()
	if err != nil {
		util.Log.Error("InsertUser failed", zap.Error(err))
		return 0, err
	}

	// 새로운 프로필을 등록합니다.
	_, err = repo.insertProfile(int(userID))
	if err != nil {
		util.Log.Error("InsertUser failed", zap.Error(err))
		return 0, err
	}

	return int(userID), nil
}

// insertProfile 는 새로운 프로필을 데이터베이스에 등록하고 새 프로필 ID를 반환합니다.
func (repo *MysqlRepo) insertProfile(userId int) (int, error) {
	query := `
		INSERT INTO profiles (user_id, bio, image, created_at, updated_at) 
		VALUES (:user_id, :bio, :image, :created_at, :updated_at)
	`
	result, err := repo.DB.NamedExec(query, table.ProfileRow{
		UserID:    userId,
		Bio:       "",
		Image:     "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		util.Log.Error("InsertProfile failed", zap.Error(err))
		return 0, err
	}

	// 새로 삽입된 프로필 ID를 가져옵니다.
	profileID, err := result.LastInsertId()
	if err != nil {
		util.Log.Error("InsertProfile failed", zap.Error(err))
		return 0, err
	}

	return int(profileID), nil
}

// FindUserByID 는 주어진 사용자 ID에 해당하는 사용자를 반환합니다.
func (repo *MysqlRepo) FindUserByID(userID int) (domain.User, error) {
	var row table.UserRow
	query := "SELECT username, email FROM users WHERE id = ?"
	err := repo.DB.Get(&row, query, userID)
	if err != nil {
		util.Log.Error("FindUserByID failed", zap.Error(err))
		return domain.User{}, err
	}
	return toUser(row), nil
}

// FindProfileByID 는 주어진 프로필 ID에 해당하는 프로필을 반환합니다.
func (repo *MysqlRepo) FindProfileByID(profileID int) (domain.Profile, error) {
	var row table.ProfileRow
	query := "SELECT bio, image FROM profiles WHERE id = ?"
	err := repo.DB.Get(&row, query, profileID)
	if err != nil {
		util.Log.Error("FindProfileByID failed", zap.Error(err))
		return domain.Profile{}, err
	}
	return toProfile(row), nil
}

// UpdateUser 는 주어진 사용자 ID에 해당하는 사용자 정보를 업데이트합니다.
func (repo *MysqlRepo) UpdateUser(userId int, user domain.User) error {
	query := `
    UPDATE users
    SET 
        username = IF(:username <> '', :username, username),
        email = IF(:email <> '', :email, email),
        password = IF(:password <> '', :password, password)
    WHERE id = :id;`

	_, err := repo.DB.NamedExec(query, table.UserRow{
		ID:       userId,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		util.Log.Error("UpdateUser failed", zap.Error(err))
		return err
	}
	return nil
}

// UpdateProfile 는 주어진 프로필 ID에 해당하는 프로필 정보를 업데이트합니다.
func (repo *MysqlRepo) UpdateProfile(profileId int, profile domain.Profile) error {
	query := `
    UPDATE profiles
    SET 
        bio = IF(:bio <> '', :bio, bio),
        image = IF(:image <> '', :image, image)
    WHERE id = :id;`

	_, err := repo.DB.NamedExec(query, table.ProfileRow{
		ID:    profileId,
		Bio:   profile.Bio,
		Image: profile.Image,
	})
	if err != nil {
		util.Log.Error("UpdateProfile failed", zap.Error(err))
		return err
	}
	return nil
}
