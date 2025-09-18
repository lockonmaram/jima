package controller

import (
	"errors"
	"fmt"
	"jima/config"
	api_entity "jima/entity/api"
	"jima/helper"
	"jima/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Authenticate(c *gin.Context)
	Register(c *gin.Context)
	SetPassword(c *gin.Context)
	ForgotPassword(c *gin.Context)
	ResetPasswordPage(c *gin.Context)
}

type authController struct {
	config      config.Config
	authService service.AuthService
}

func NewAuthController(config config.Config, authService service.AuthService) AuthController {
	return &authController{
		config,
		authService,
	}
}

func (ac *authController) Authenticate(c *gin.Context) {
	request := api_entity.AuthAuthenticateRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	response, err := ac.authService.Authenticate(c, request.UserParam, request.Password)
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

	response, err := ac.authService.Register(c, request)
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

func (ac *authController) SetPassword(c *gin.Context) {
	request := api_entity.AuthSetPasswordRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	err := ac.authService.SetPassword(c, request)
	if err != nil {
		helper.HandleResponse(c, helper.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	helper.HandleResponse(c, helper.Response{
		Status: http.StatusOK,
	})
}

func (ac *authController) ForgotPassword(c *gin.Context) {
	request := api_entity.AuthForgotPasswordRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	err := ac.authService.ForgotPassword(c, request)
	if err != nil {
		helper.HandleResponse(c, helper.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	helper.HandleResponse(c, helper.Response{
		Status:  http.StatusOK,
		Message: "reset password link is sent to your email",
	})
}

func (ac *authController) ResetPasswordPage(c *gin.Context) {
	request := api_entity.AuthResetPasswordPageRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	setPasswordURL := fmt.Sprintf("%s:%d/api/v1/auth/set-password", ac.config.BaseURL, ac.config.Port)

	c.HTML(http.StatusOK, "reset-password-form.template.html", gin.H{
		"Action": setPasswordURL,
	})
}
