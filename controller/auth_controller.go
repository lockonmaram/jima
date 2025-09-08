package controller

import (
	"jima/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Authenticate(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (ac *authController) Authenticate(c *gin.Context) {}
func (ac *authController) Register(c *gin.Context)     {}
