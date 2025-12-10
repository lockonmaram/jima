package api_entity

type UsersCreateUserRequest struct {
	Username string `json:"username" validation:"required"`
	Password string `json:"password" validation:"required"`
	Email    string `json:"email" validation:"required,email"`
	Name     string `json:"name" validation:"required"`
	Role     string `json:"role" validation:"required"`
}
type UsersCreateUserResponse struct {
	Serial   string `json:"serial"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type UsersUpdateUserProfileRequest struct {
	Serial string `uri:"serial" validation:"required"`
	Name   string `json:"name"`
}
type UsersUpdateUserProfileResponse struct {
	Name string `json:"name,omitempty"`
}

type UsersChangePasswordRequest struct {
	Serial   string `uri:"serial" validation:"required"`
	Password string `json:"password"`
}
