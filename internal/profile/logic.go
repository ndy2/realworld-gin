package profile

import (
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/util"
)

type Logic struct {
	repo Repo
}

// NewLogic creates and returns a new Logic.
func NewLogic(repo Repo) Logic {
	return Logic{repo: repo}
}

// GetProfile returns the profile of the given user.
func (l Logic) GetProfile(currentUsername string, profileId int) (GetProfileResponse, error) {
	// Get the profile of the given user.
	profile, err := l.repo.FindProfile(profileId)
	if err != nil {
		util.Log.Error("FindProfile failed", zap.Error(err))
		return GetProfileResponse{}, err
	}

	return GetProfileResponse{
		Username:  currentUsername,
		Bio:       profile.Bio,
		Image:     profile.Image,
		Following: false,
	}, nil
}
