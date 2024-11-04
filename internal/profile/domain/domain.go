package domain

type Profile struct {
	Username string
	Bio      string
	Image    string
}

type Following bool

type Repo interface {
	// FindProfile returns the profile of the given user.
	FindProfile(profileID int) (Profile, error)

	// GetProfileWithFollowing returns the profile of the given user and whether the current user is following them.
	//GetProfileWithFollowing(profileID int) (Profile, Following, error)
}
