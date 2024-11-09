package util

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func init() {
	// Custom validators 등록
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("username", validateUsername)
		_ = v.RegisterValidation("email", validateEmail)
		_ = v.RegisterValidation("password", validatePassword)
		_ = v.RegisterValidation("image", validateImage)
		_ = v.RegisterValidation("bio", validateBio)
	}
}

// Username 은 20자 이하의 문자열이어야 함
func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return len(username) <= 20
}

// Email 은 256자 이하의 이메일 형식이어야 함
var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return len(email) <= 256 && emailRegex.MatchString(email)
}

// Password 는 최소 8자, 최대 16자 대,소문자와 특수문자를 포함해야 함
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var hasGoodLen, hasUpper, hasLower, hasDigit, hasSpecial bool

	if len(password) >= 8 && len(password) <= 16 {
		hasGoodLen = true
	}
	for _, c := range password {
		switch {
		case 'A' <= c && c <= 'Z':
			hasUpper = true
		case 'a' <= c && c <= 'z':
			hasLower = true
		case '0' <= c && c <= '9':
			hasDigit = true
		case c == '!' || c == '@' || c == '#' || c == '$' || c == '%' || c == '^' || c == '&':
			hasSpecial = true
		}
	}

	return hasGoodLen && hasUpper && hasLower && hasDigit && hasSpecial
}

// Image 는 256자 이하의 URL 형식이어야 함
var imageRegex = regexp.MustCompile(`^https?://.+$`)

func validateImage(fl validator.FieldLevel) bool {
	image := fl.Field().String()
	return len(image) <= 256 && imageRegex.MatchString(image)
}

// Bio 는 256자 이하의 문자열이어야 함
func validateBio(fl validator.FieldLevel) bool {
	bio := fl.Field().String()
	return len(bio) <= 256
}
