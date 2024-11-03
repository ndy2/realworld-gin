package app

import "errors"

var ErrPasswordMismatch = errors.New("password mismatch")

var ErrUserNotFound = errors.New("user not found")
var ErrProfileNotFound = errors.New("profile not found")
