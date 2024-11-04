package domain

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
	//
	// Parameters:
	// - email: the email to check
	//
	// Returns:
	// - bool: true if the user exists, false otherwise
	// - error: an error if the operation fails
	CheckUserExists(email string) (bool, error)

	// InsertUser insert a new user
	//
	// Parameters:
	// - u: the user to insert
	//
	// Returns:
	// - int: the ID of the new user
	// - error: an error if the operation fails
	InsertUser(u User) (int, error)

	// FindUserByID finds a user by ID
	//
	// Parameters:
	// - userID: the ID of the user to find
	//
	// Returns:
	// - User: the user found
	// - error: an error if the operation fails (e.g. SQLConnectionError)
	//			or the user is not found (e.g. SQLNoRows)
	FindUserByID(userID int) (User, error)

	// FindProfileByID finds a profile by ID
	//
	// Parameters:
	// - profileID: the ID of the profile to find
	//
	// Returns:
	// - Profile: the profile found
	// - error: an error if the operation fails (e.g. SQLConnectionError)
	//			or the profile is not found (e.g. SQLNoRows)
	FindProfileByID(profileID int) (Profile, error)

	// UpdateUser updates a user
	//
	// Parameters:
	// - userId: the ID of the user to update
	// - user: the user data to update
	//
	// Returns:
	// - error: an error if the operation fails
	UpdateUser(userId int, user User) error

	// UpdateProfile updates a profile
	//
	// Parameters:
	// - profileId: the ID of the profile to update
	// - profile: the profile data to update
	//
	// Returns:
	// - error: an error if the operation fails
	UpdateProfile(profileId int, profile Profile) error
}
