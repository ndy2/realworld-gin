package user

type User struct {
	Username string
	Email    string
	Password string
}

type Repo interface {
	// CheckUserExists checks if a user with the given email exists
	CheckUserExists(email string) (bool, error)

	// InsertUser insert a new user
	InsertUser(u User) (int, error)
}
