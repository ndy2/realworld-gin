package app

import (
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/profile/domain"
	"ndy/realworld-gin/internal/profile/dto"
	"ndy/realworld-gin/internal/util"
)

type Logic struct {
	repo domain.Repo
}

// NewLogic creates and returns a new Logic.
func NewLogic(repo domain.Repo) Logic {
	return Logic{repo: repo}
}

// GetProfile returns the profile of the given user.
func (l Logic) GetProfile(currentUsername string, profileId int) (dto.GetProfileResponse, error) {
	// Get the profile of the given user.
	profile, err := l.repo.FindProfile(profileId)
	if err != nil {
		util.Log.Error("FindProfile failed", zap.Error(err))
		return dto.GetProfileResponse{}, err
	}

	return dto.GetProfileResponse{
		Username:  currentUsername,
		Bio:       profile.Bio,
		Image:     profile.Image,
		Following: false,
	}, nil
}
