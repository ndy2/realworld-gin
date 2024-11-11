package app

import (
	"context"
	"ndy/realworld-gin/internal/user/dto"
)

type Logic interface {
	Register(username, email, password string) (int, error)
	GetCurrentUser(userID, profileId int) (dto.GetCurrentUserResponse, error)
	UpdateUser(ctx context.Context, email, username, password, image, bio string) (dto.UpdateUserResponse, error)
}
