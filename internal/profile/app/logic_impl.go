package app

import (
	"go.uber.org/zap"
	"ndy/realworld-gin/internal/profile/domain"
	"ndy/realworld-gin/internal/profile/dto"
	"ndy/realworld-gin/internal/util"
)

type LogicImpl struct {
	repo domain.Repo
}

// NewLogicImpl creates and returns a new Logic.
func NewLogicImpl(repo domain.Repo) Logic {
	return LogicImpl{repo: repo}
}

// GetProfile returns the profile of the given user.
func (l LogicImpl) GetProfile(currentUserId, currentUserProfileId int, currentUsername, targetUsername string) (dto.GetProfileResponse, error) {
	var profile domain.Profile
	var following domain.Following = false
	var err error

	if currentUsername == targetUsername {
		// Authenticated user is the same as the target user.
		profile, err = l.repo.FindProfile(currentUserProfileId)
		if err != nil {
			util.Log.Error("FindProfile failed", zap.Error(err))
			return dto.GetProfileResponse{}, err
		}
	} else if currentUsername == "" {
		// Unauthenticated user is viewing the profile of another user.
		profile, err = l.repo.FindProfileByUsername(targetUsername)
		if err != nil {
			util.Log.Error("FindProfileByUsername failed", zap.Error(err))
			return dto.GetProfileResponse{}, err
		}
	} else {
		// Authenticated user is viewing the profile of another user.
		profile, following, err = l.repo.FindProfileWithFollowingByUsername(targetUsername, currentUserId)
	}

	return dto.GetProfileResponse{
		Username:  targetUsername,
		Bio:       profile.Bio,
		Image:     profile.Image,
		Following: bool(following),
	}, nil
}
