package controller

import (
	"errors"
	api_entity "jima/entity/api"
	"jima/helper"
	"jima/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(c *gin.Context)
	UpdateUserProfile(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService,
	}
}

func (uc *userController) CreateUser(c *gin.Context) {
	request := api_entity.UserCreateUserRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	response, err := uc.userService.CreateUser(c, request)
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

func (uc *userController) UpdateUserProfile(c *gin.Context) {
	request := api_entity.UserUpdateUserProfileRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	response, err := uc.userService.UpdateUserProfile(c, request)
	if err != nil {
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

func (uc *userController) ChangePassword(c *gin.Context) {
	request := api_entity.UserChangePasswordRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	err := uc.userService.ChangePassword(c, request)
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
	})
}
