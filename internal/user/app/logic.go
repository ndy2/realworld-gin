package app

import (
	"context"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
	"ndy/realworld-gin/internal/user/domain"
	"ndy/realworld-gin/internal/user/dto"
	"ndy/realworld-gin/internal/util"
)

type Logic interface {
	Register(username, email, password string) (int, error)
	GetCurrentUser(userID, profileId int) (dto.GetCurrentUserResponse, error)
	UpdateUser(ctx context.Context, email, username, password, image, bio string) (dto.UpdateUserResponse, error)
}

type LogicImpl struct {
	repo domain.Repo
}

// NewLogicImpl 는 새로운 LogicImpl 을 생성하고 반환합니다.
func NewLogicImpl(repo domain.Repo) Logic {
	return LogicImpl{repo: repo}
}

// Register 는 새로운 사용자를 등록합니다.
func (l LogicImpl) Register(
	username string,
	email string,
	password string,
) (int, error) {
	// 사용자가 이미 존재하는지 확인합니다.
	exists, err := l.repo.CheckUserExists(email)
	if err != nil {
		util.Log.Error("CheckUserExists failed", zap.Error(err))
		return 0, err
	}
	if exists {
		util.Log.Info("Email already registered", zap.String("email", email))
		return 0, EmailAlreadyRegistered
	}

	// 비밀번호를 해시화합니다.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		util.Log.Error("GenerateFromPassword failed", zap.Error(err))
		return 0, err
	}
	// 사용자를 등록합니다.
	id, err := l.repo.InsertUser(domain.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	})
	if err != nil {
		util.Log.Error("InsertUser failed", zap.Error(err))
		return 0, err
	}

	// 새로 등록된 사용자 ID를 반환합니다.
	return id, nil
}

// GetCurrentUser 는 현재 사용자 정보를 반환합니다.
func (l LogicImpl) GetCurrentUser(userID, profileId int) (dto.GetCurrentUserResponse, error) {
	var user domain.User
	var profile domain.Profile

	g, _ := errgroup.WithContext(context.Background())

	// 사용자 정보를 조회합니다.
	g.Go(func() error {
		u, err := l.repo.FindUserByID(userID)
		if err != nil {
			return err
		}
		user = u
		return nil
	})

	// 프로필 정보를 조회합니다.
	g.Go(func() error {
		p, err := l.repo.FindProfileByID(profileId)
		if err != nil {
			return err
		}
		profile = p
		return nil
	})

	// 에러가 발생하면 에러를 반환합니다.
	if err := g.Wait(); err != nil {
		return dto.GetCurrentUserResponse{}, err
	}

	// 사용자 정보와 프로필 정보를 반환합니다.
	return dto.GetCurrentUserResponse{
		Username: user.Username,
		Email:    user.Email,
		Bio:      profile.Bio,
		Image:    profile.Image,
		Token:    "",
	}, nil
}

// UpdateUser 는 사용자 정보를 업데이트합니다.
func (l LogicImpl) UpdateUser(ctx context.Context, email, username, password, image, bio string) (dto.UpdateUserResponse, error) {
	// 사용자 ID를 context 에서 추출합니다.
	userId, _ := ctx.Value("userId").(int)
	profileId, _ := ctx.Value("profileId").(int)

	// 비밀번호를 해시화합니다.
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return dto.UpdateUserResponse{}, err
		}
		password = string(hashedPassword)
	}

	// 사용자 정보를 업데이트합니다.
	err := l.repo.UpdateUser(userId, domain.User{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		util.Log.Error("UpdateUser failed", zap.Error(err))
		return dto.UpdateUserResponse{}, err
	}

	// 프로필 정보를 업데이트합니다.
	err = l.repo.UpdateProfile(profileId, domain.Profile{
		Bio:   bio,
		Image: image,
	})
	if err != nil {
		util.Log.Error("UpdateProfile failed", zap.Error(err))
		return dto.UpdateUserResponse{}, err
	}

	// 사용자 정보를 조회합니다.
	updatedUser, err := l.repo.FindUserByID(userId)
	if err != nil {
		util.Log.Error("FindUserByID failed", zap.Error(err))
		return dto.UpdateUserResponse{}, err
	}

	// 프로필 정보를 조회합니다.
	updatedProfile, err := l.repo.FindProfileByID(profileId)
	if err != nil {
		util.Log.Error("FindProfileByID failed", zap.Error(err))
		return dto.UpdateUserResponse{}, err
	}

	// 업데이트된 사용자 정보를 반환합니다.
	return dto.UpdateUserResponse{
		Username: updatedUser.Username,
		Email:    updatedUser.Email,
		Bio:      updatedProfile.Bio,
		Image:    updatedProfile.Image,
		Token:    "",
	}, nil
}
