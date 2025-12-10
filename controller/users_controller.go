package controller

import (
	"errors"
	api_entity "jima/entity/api"
	"jima/helper"
	"jima/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController interface {
	CreateUser(c *gin.Context)
	UpdateUserProfile(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type usersController struct {
	usersService service.UsersService
}

func NewUsersController(usersService service.UsersService) UsersController {
	return &usersController{
		usersService,
	}
}

func (uc *usersController) CreateUser(c *gin.Context) {
	request := api_entity.UsersCreateUserRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	response, err := uc.usersService.CreateUser(c, request)
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

func (uc *usersController) UpdateUserProfile(c *gin.Context) {
	request := api_entity.UsersUpdateUserProfileRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	response, err := uc.usersService.UpdateUserProfile(c, request)
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

func (uc *usersController) ChangePassword(c *gin.Context) {
	request := api_entity.UsersChangePasswordRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	err := uc.usersService.ChangePassword(c, request)
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
