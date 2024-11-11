package app

import (
	"ndy/realworld-gin/internal/auth/dto"
)

type Logic interface {
	Login(email string, password string) (dto.LoginResponse, error)
}
