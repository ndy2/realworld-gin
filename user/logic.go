package user

import (
	"context"
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

// GetCurrentUser 는 현재 사용자 정보를 반환합니다.
func (l Logic) GetCurrentUser(userID, profileId int) (GetCurrentUserResponse, error) {
	// 사용자 정보를 조회합니다.
	user, err := l.repo.FindUserByID(userID)
	if err != nil {
		return GetCurrentUserResponse{}, err
	}

	// 프로필 정보를 조회합니다.
	profile, err := l.repo.FindProfileByID(profileId)
	if err != nil {
		return GetCurrentUserResponse{}, err
	}

	// 사용자 정보와 프로필 정보를 반환합니다.
	return GetCurrentUserResponse{
		Username: user.Username,
		Email:    user.Email,
		Bio:      profile.Bio,
		Image:    profile.Image,
		Token:    "",
	}, nil
}

// UpdateUser 는 사용자 정보를 업데이트합니다.
func (l Logic) UpdateUser(ctx context.Context, email, username, password, image, bio string) (UpdateUserResponse, error) {
	// 사용자 ID를 context 에서 추출합니다.
	userId, _ := ctx.Value("userId").(int)
	profileId, _ := ctx.Value("profileId").(int)

	// 비밀번호를 해시화합니다.
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return UpdateUserResponse{}, err
		}
		password = string(hashedPassword)
	}

	// 사용자 정보를 업데이트합니다.
	err := l.repo.UpdateUser(userId, User{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return UpdateUserResponse{}, err
	}

	// 프로필 정보를 업데이트합니다.
	err = l.repo.UpdateProfile(profileId, Profile{
		Bio:   bio,
		Image: image,
	})
	if err != nil {
		return UpdateUserResponse{}, err
	}

	// 사용자 정보를 조회합니다.
	updatedUser, err := l.repo.FindUserByID(userId)
	if err != nil {
		return UpdateUserResponse{}, err
	}

	// 프로필 정보를 조회합니다.
	updatedProfile, err := l.repo.FindProfileByID(profileId)
	if err != nil {
		return UpdateUserResponse{}, err
	}

	// 업데이트된 사용자 정보를 반환합니다.
	return UpdateUserResponse{
		Username: updatedUser.Username,
		Email:    updatedUser.Email,
		Bio:      updatedProfile.Bio,
		Image:    updatedProfile.Image,
		Token:    "",
	}, nil
}
