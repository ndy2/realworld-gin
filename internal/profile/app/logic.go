package app

import (
	"ndy/realworld-gin/internal/profile/dto"
)

type Logic interface {
	GetProfile(currentUserId, currentUserProfileId int, currentUsername, targetUsername string) (dto.GetProfileResponse, error)
}
