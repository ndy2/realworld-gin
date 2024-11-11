package infra

import (
	"ndy/realworld-gin/internal/profile/domain"
	"ndy/realworld-gin/internal/util/table"
)

func toProfile(row table.ProfileRow) domain.Profile {
	return domain.Profile{
		Bio:   row.Bio,
		Image: row.Image,
	}
}
