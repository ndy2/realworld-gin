package infra

import (
	"ndy/realworld-gin/internal/auth/domain"
	"ndy/realworld-gin/internal/util/table"
)

func toUser(row table.UserRow) domain.User {
	return domain.User{
		Id:       row.ID,
		Username: row.Username,
		Email:    row.Email,
		Password: row.Password,
	}
}

func toProfile(row table.ProfileRow) domain.Profile {
	return domain.Profile{
		Id:     row.ID,
		UserID: row.UserID,
		Bio:    row.Bio,
		Image:  row.Image,
	}
}
