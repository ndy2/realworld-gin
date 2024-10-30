package auth

type User struct {
	Email    string
	Password string
}

type Repo interface {
	// FindUserByEmail checks if a user with the given email exists
	FindUserByEmail(email string) (User, error)
}
