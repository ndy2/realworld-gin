package auth

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"ndy/realworld-gin/logger"
)

type Logic interface {
	Login(email string, password string) (LoginResponse, error)
}

type LogicImpl struct {
	repo Repo
}

// NewLogicImpl 는 새로운 LogicImpl 을 생성하고 반환합니다.
func NewLogicImpl(repo Repo) Logic {
	return LogicImpl{repo: repo}
}

// Login 는 사용자를 인증하고 토큰을 반환합니다.
func (l LogicImpl) Login(email string, password string) (LoginResponse, error) {
	// 사용자가 존재하는지 확인합니다.
	user, err := l.repo.FindUserByEmail(email)
	if errors.Is(err, ErrUserNotFound) {
		logger.Log.Info("User not found", zap.String("email", email))
		return LoginResponse{}, ErrUserNotFound
	}
	if err != nil {
		logger.Log.Error("FindUserByEmail failed", zap.Error(err))
		return LoginResponse{}, err
	}

	// 비밀번호가 일치하는지 확인합니다.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrPasswordMismatch, err)
		logger.Log.Info("CompareHashAndPassword failed", zap.Error(err))
		return LoginResponse{}, err
	}

	// 프로필 정보를 조회합니다.
	profile, err := l.repo.FindProfileByUserID(user.Id)
	if err != nil {
		logger.Log.Error("FindProfileByUserID failed", zap.Error(err))
		return LoginResponse{}, err
	}

	// 토큰을 생성합니다.
	token, err := generate(user, profile)
	if err != nil {
		logger.Log.Error("generate failed", zap.Error(err))
		return LoginResponse{}, err
	}

	// 사용자 정보와 토큰을 반환합니다.
	return LoginResponse{
		Email:    user.Email,
		Token:    token,
		Username: profile.Username,
		Bio:      profile.Bio,
		Image:    profile.Image,
	}, nil
}
