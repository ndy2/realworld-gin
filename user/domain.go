package user

type User struct {
	Username string
	Email    string
	Password string
}

type Profile struct {
	Bio   string
	Image string
}

type Repo interface {
	// CheckUserExists checks if a user with the given email exists
	CheckUserExists(email string) (bool, error)

	// InsertUser insert a new user
	InsertUser(u User) (int, error)

	// FindUserByID finds a user by ID
	FindUserByID(userID int) (User, error)

	// FindProfileByID finds a profile by ID
	FindProfileByID(profileID int) (Profile, error)

	// UpdateUser updates a user
	UpdateUser(id int, user User) error

	// UpdateProfile updates a profile
	UpdateProfile(id int, profile Profile) error
}
