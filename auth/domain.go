package auth

type User struct {
	Id       int
	Email    string
	Password string
}

type Profile struct {
	Id       int
	Username string
	UserID   int
	Bio      string
	Image    string
}

type Repo interface {
	// FindUserByEmail checks if a user with the given email exists
	FindUserByEmail(email string) (User, error)

	// FindProfileByUserID finds a profile by user ID
	FindProfileByUserID(userID int) (Profile, error)
}
