package dto

type RegistrationRequest struct {
	Username string `json:"username" binding:"username"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"password"`
}
type RegistrationResponse userResponse

type GetCurrentUserResponse userResponse

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"email"`
	Username string `json:"username" binding:"username"`
	Password string `json:"password" binding:"password"`
	Image    string `json:"image" binding:"image"`
	Bio      string `json:"bio" binding:"bio"`
}
type UpdateUserResponse userResponse

type userResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
