package auth

import (
	"testing"
)

func TestUserStruct(t *testing.T) {
	// given
	user := User{
		Id:       1,
		Email:    "test@example.com",
		Password: "password123",
	}

	// when & then
	if user.Id != 1 {
		t.Errorf("expected Id to be 1, got %d", user.Id)
	}
	if user.Email != "test@example.com" {
		t.Errorf("expected Email to be 'test@example.com', got %s", user.Email)
	}
	if user.Password != "password123" {
		t.Errorf("expected Password to be 'password123', got %s", user.Password)
	}
}
