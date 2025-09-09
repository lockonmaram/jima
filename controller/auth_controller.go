package controller

import (
	"errors"
	api_entity "jima/entity/api"
	"jima/helper"
	"jima/service"
	"net/http"

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
		authService,
	}
}

func (ac *authController) Authenticate(c *gin.Context) {
	request := api_entity.AuthAuthenticateRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	response, err := ac.authService.Authenticate(request.UserParam, request.Password)
	if err != nil {
		if errors.Is(err, helper.ErrUserNotFound) {
			helper.HandleResponse(c, helper.Response{
				Status:  http.StatusUnauthorized,
				Message: helper.ErrInvalidUsernameEmail.Error(),
			})
			return
		} else if errors.Is(err, helper.ErrInvalidPassword) {
			helper.HandleResponse(c, helper.Response{
				Status:  http.StatusUnauthorized,
				Message: helper.ErrInvalidPassword.Error(),
			})
			return
		}

		helper.HandleResponse(c, helper.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	helper.HandleResponse(c, helper.Response{
		Status: http.StatusOK,
		Data:   response,
	})
}

func (ac *authController) Register(c *gin.Context) {
	request := api_entity.AuthRegisterRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	response, err := ac.authService.Register(request)
	if err != nil {
		if errors.Is(err, helper.ErrUserAlreadyExists) {
			helper.HandleResponse(c, helper.Response{
				Status: http.StatusConflict,
				Error:  err.Error(),
			})
			return
		}

		helper.HandleResponse(c, helper.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	helper.HandleResponse(c, helper.Response{
		Status: http.StatusOK,
		Data:   response,
	})
}
