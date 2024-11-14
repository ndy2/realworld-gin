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

type UserProfile struct {
	User
	Profile
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

	// InsertUserProfile inserts a new user
	//
	// Parameters:
	// - up: the user profile to insert
	//
	// Returns:
	InsertUserProfile(up UserProfile) (int, error)

	// FindUserProfileByID finds a user profile by user ID and profile ID
	//
	// Parameters:
	// - userID: the ID of the user
	// - profileID: the ID of the profile
	//
	// Returns:
	// - UserProfile: the user profile
	FindUserProfileByID(userID, profileID int) (UserProfile, error)

	// UpdateUserProfile updates a user profile
	//
	// Parameters:
	// - userId: the ID of the user to update
	// - profileId: the ID of the profile to update
	// - up: the new user profile
	UpdateUserProfile(userId, profileId int, up UserProfile) error
}
