package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordRegex(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		isValid  bool
	}{
		{"Valid password", "Password1!", true},
		{"Missing uppercase letter", "password1!", false},
		{"Missing digit", "Password!", false},
		{"Missing special character", "Password1", false},
		{"Too short", "P1!", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := passwordRegex.MatchString(tc.password)
			assert.Equal(t, tc.isValid, result, "Password validation failed for: %s", tc.name)
		})
	}
}
