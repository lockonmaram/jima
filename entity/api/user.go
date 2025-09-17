package api_entity

type UserCreateUserRequest struct {
	Username string `json:"username" validation:"required"`
	Password string `json:"password" validation:"required"`
	Email    string `json:"email" validation:"required,email"`
	Name     string `json:"name" validation:"required"`
	Role     string `json:"role" validation:"required"`
}
type UserCreateUserResponse struct {
	Serial   string `json:"serial"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type UserUpdateUserProfileRequest struct {
	Serial string `uri:"serial" validation:"required"`
	Name   string `json:"name"`
}
type UserUpdateUserProfileResponse struct {
	Name string `json:"name,omitempty"`
}

type UserChangePasswordRequest struct {
	Serial   string `uri:"serial" validation:"required"`
	Password string `json:"password"`
}
