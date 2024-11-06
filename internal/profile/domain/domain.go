package domain

type Profile struct {
	Bio   string
	Image string
}

type Following bool

type Repo interface {
	// FindProfile returns the profile of the given user.
	FindProfile(profileID int) (Profile, error)

	// FindProfileByUsername returns the profile of the given user.
	FindProfileByUsername(username string) (Profile, error)

	// FindProfileWithFollowingByUsername returns the profile of the given user and whether the current user is following the target user.
	FindProfileWithFollowingByUsername(username string, currentUserId int) (Profile, Following, error)
}
