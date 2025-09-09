package api_entity

type AuthAuthenticateRequest struct {
	Username string `json:"username" validation:"required"`
	Password string `json:"password" validation:"required"`
}
type AuthAuthenticateResponse struct {
	Username string `json:"username" validation:"required"`
	Password string `json:"password" validation:"required"`
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
