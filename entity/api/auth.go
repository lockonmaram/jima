package api_entity

type AuthAuthenticateRequest struct {
	UserParam string `json:"userParam" validation:"required"` // username or email
	Password  string `json:"password" validation:"required"`
}
type AuthAuthenticateResponse struct {
	Token    string `json:"token"`
	Serial   string `json:"serial"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type AuthRegisterRequest struct {
	Username string `json:"username" validation:"required"`
	Password string `json:"password" validation:"required"`
	Email    string `json:"email" validation:"required,email"`
	Name     string `json:"name" validation:"required"`
	Role     string `json:"role" validation:"required"`
}
type AuthRegisterResponse struct {
	Serial   string `json:"serial"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}
