package api_entity

type UserCreateRequest struct {
	Username string `json:"username" validation:"required"`
	Password string `json:"password" validation:"required"`
	Email    string `json:"email" validation:"required,email"`
	Name     string `json:"name" validation:"required"`
	Role     string `json:"role" validation:"required"`
}
type UserCreateResponse struct {
	Serial   string `json:"serial"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}
