package user

// Registration
type registrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type registrationResponse userResponse

// GetCurrentUser
type getCurrentUserResponse userResponse

// UpdateUser
type updateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Image    string `json:"image"`
	Bio      string `json:"bio"`
}
type updateUserResponse userResponse

type userResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
