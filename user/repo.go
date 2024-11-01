package user

import (
	"database/sql"
	"go.uber.org/zap"
	"ndy/realworld-gin/logger"
	"time"
)

type MysqlRepo struct {
	db *sql.DB
}

// NewMysqlRepo 는 MysqlRepo 를 생성하고 반환합니다.
func NewMysqlRepo(db *sql.DB) *MysqlRepo {
	return &MysqlRepo{db: db}
}

// CheckUserExists 는 주어진 이메일을 가진 사용자가 있는지 확인합니다.
func (repo *MysqlRepo) CheckUserExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)"
	err := repo.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		logger.Error("CheckUserExists failed", zap.Error(err))
		return false, err
	}
	return exists, nil
}

// InsertUser 는 새로운 사용자를 데이터베이스에 등록하고 새 사용자 ID를 반환합니다.
func (repo *MysqlRepo) InsertUser(u User) (int, error) {
	query := "INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, u.Username, u.Email, u.Password, time.Now(), time.Now())
	if err != nil {
		logger.Error("InsertUser failed", zap.Error(err))
		return 0, err
	}

	// 새로 삽입된 사용자 ID를 가져옵니다.
	userID, err := result.LastInsertId()
	if err != nil {
		logger.Error("InsertUser failed", zap.Error(err))
		return 0, err
	}

	// 새로운 프로필을 등록합니다.
	_, err = repo.insertProfile(int(userID))
	if err != nil {
		logger.Error("InsertUser failed", zap.Error(err))
		return 0, err
	}

	return int(userID), nil
}

// insertProfile 는 새로운 프로필을 데이터베이스에 등록하고 새 프로필 ID를 반환합니다.
func (repo *MysqlRepo) insertProfile(userId int) (int, error) {
	query := "INSERT INTO profiles (user_id, bio, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, userId, "", "", time.Now(), time.Now())
	if err != nil {
		logger.Error("InsertProfile failed", zap.Error(err))
		return 0, err
	}

	// 새로 삽입된 프로필 ID를 가져옵니다.
	profileID, err := result.LastInsertId()
	if err != nil {
		logger.Error("InsertProfile failed", zap.Error(err))
		return 0, err
	}

	return int(profileID), nil
}

// FindUserByID 는 주어진 사용자 ID에 해당하는 사용자를 반환합니다.
func (repo *MysqlRepo) FindUserByID(userID int) (User, error) {
	var user User
	query := "SELECT username, email FROM users WHERE id = ?"
	err := repo.db.QueryRow(query, userID).Scan(&user.Username, &user.Email)
	if err != nil {
		logger.Error("FindUserByID failed", zap.Error(err))
		return User{}, err
	}
	return user, nil
}

// FindProfileByID 는 주어진 프로필 ID에 해당하는 프로필을 반환합니다.
func (repo *MysqlRepo) FindProfileByID(profileID int) (Profile, error) {
	var profile Profile
	query := "SELECT bio, image FROM profiles WHERE id = ?"
	err := repo.db.QueryRow(query, profileID).Scan(&profile.Bio, &profile.Image)
	if err != nil {
		logger.Error("FindProfileByID failed", zap.Error(err))
		return Profile{}, err
	}
	return profile, nil
}

// UpdateUser 는 주어진 사용자 ID에 해당하는 사용자 정보를 업데이트합니다.
func (repo *MysqlRepo) UpdateUser(userId int, user User) error {
	query := `
    UPDATE users
    SET 
        username = IF(? <> '', ?, username),
        email = IF(? <> '', ?, email),
        password = IF(? <> '', ?, password)
    WHERE id = ?;`

	_, err := repo.db.Exec(query, user.Username, user.Username, user.Email, user.Email, user.Password, user.Password, userId)
	if err != nil {
		logger.Error("UpdateUser failed", zap.Error(err))
		return err
	}
	return nil
}

// UpdateProfile 는 주어진 프로필 ID에 해당하는 프로필 정보를 업데이트합니다.
func (repo *MysqlRepo) UpdateProfile(profileId int, profile Profile) error {
	query := `
    UPDATE profiles
    SET 
        bio = IF(? <> '', ?, bio),
        image = IF(? <> '', ?, image)
    WHERE id = ?;`
	_, err := repo.db.Exec(query, profile.Bio, profile.Bio, profile.Image, profile.Image, profileId)
	if err != nil {
		logger.Error("UpdateProfile failed", zap.Error(err))
		return err
	}
	return nil
}
