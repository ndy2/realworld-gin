package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_validateUsername(t *testing.T) {
	runValidationTests(t, "username", validateUsername, []validatorTest{
		{"valid username", "username", true},
		{"too long username", "too long username username username username", false},
	})
}

func Test_validateEmail(t *testing.T) {
	runValidationTests(t, "email", validateEmail, []validatorTest{
		{"valid email", "jake@jake.jake", true},
		{"invalid email", "jake@jake", false},
	})
}

func Test_validatePassword(t *testing.T) {
	runValidationTests(t, "password", validatePassword, []validatorTest{
		{"valid password", "Password1!", true},
		{"too short password", "Pass1!", false},
		{"too long password", "Password12345678901!", false},
		{"no uppercase letter", "password1!", false},
		{"no lowercase letter", "PASSWORD1!", false},
		{"no digit", "Password!", false},
		{"no special character", "Password1", false},
	})
}

func Test_validateImage(t *testing.T) {
	runValidationTests(t, "image", validateImage, []validatorTest{
		{"valid image", "https://image.com/image.jpg", true},
		{"invalid image", "image.jpg", false},
	})
}

func Test_validateBio(t *testing.T) {
	runValidationTests(t, "bio", validateBio, []validatorTest{
		{"valid bio", "bio", true},
		{"valid 256 characters", strings.Repeat("a", 256), true},
		{"invalid 257 characters", strings.Repeat("a", 257), false},
	})
}

type validatorTest struct {
	name  string
	value string
	valid bool
}

func runValidationTests(t *testing.T, tag string, validateFunc validator.Func, tests []validatorTest) {
	validate := validator.New()
	_ = validate.RegisterValidation(tag, validateFunc)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Var(tt.value, tag)
			assert.Equal(t, tt.valid, err == nil, "unexpected validation result for test case: %s", tt.name)
		})
	}
}
