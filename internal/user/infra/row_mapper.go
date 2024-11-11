package infra

import (
	"ndy/realworld-gin/internal/user/domain"
	"ndy/realworld-gin/internal/util/table"
)

func toUser(row table.UserRow) domain.User {
	return domain.User{
		Username: row.Username,
		Email:    row.Email,
		Password: row.Password,
	}
}

func toProfile(row table.ProfileRow) domain.Profile {
	return domain.Profile{
		Bio:   row.Bio,
		Image: row.Image,
	}
}
