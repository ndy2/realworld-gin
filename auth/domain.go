package auth

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

type Profile struct {
	Id     int
	UserID int
	Bio    string
	Image  string
}

type Repo interface {
	// FindUserByEmail checks if a user with the given email exists
	//
	// Parameters:
	//  - email: the email address to check
	//
	// Returns:
	//  - User: the user with the given email
	//  - error: nil if the user is found, otherwise:
	//      - ErrUserNotFound: the user with the given email is not found
	//      - other errors (e.g. database error)
	FindUserByEmail(email string) (User, error)

	// FindProfileByUserID finds a profile by user ID
	//
	// Parameters:
	//  - userID: the user ID to search for
	//
	// Returns:
	//  - Profile: the profile with the given user ID
	//  - error: nil if the profile is found, otherwise:
	//      - ErrProfileNotFound: the profile with the given user ID is not found
	//      - other errors (e.g. database error)
	FindProfileByUserID(userID int) (Profile, error)
}
