package dto

type RegistrationRequest struct {
	Username string `json:"username" binding:"username"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"password"`
}
type RegistrationResponse userResponse

type GetCurrentUserResponse userResponse

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Username string `json:"username" binding:"omitempty,username"`
	Password string `json:"password" binding:"omitempty,password"`
	Image    string `json:"image" binding:"omitempty,image"`
	Bio      string `json:"bio" binding:"omitempty,bio"`
}
type UpdateUserResponse userResponse

type userResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
