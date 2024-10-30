package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Logic struct {
	repo Repo
}

// NewLogic 는 새로운 Logic을 생성하고 반환합니다.
func NewLogic(repo Repo) Logic {
	return Logic{repo: repo}
}

// Register 는 새로운 사용자를 등록합니다.
func (l Logic) Register(
	username string,
	email string,
	password string,
) (int, error) {
	// 사용자가 이미 존재하는지 확인합니다.
	exists, err := l.repo.CheckUserExists(email)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, fmt.Errorf("user already exists")
	}

	// 비밀번호를 해시화합니다.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	// 사용자를 등록합니다.
	id, err := l.repo.InsertUser(User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return 0, err
	}

	// 새로 등록된 사용자 ID를 반환합니다.
	return id, nil
}
