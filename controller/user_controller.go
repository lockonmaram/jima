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
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService,
	}
}

func (ac *userController) CreateUser(c *gin.Context) {
	request := api_entity.UserCreateRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	response, err := ac.userService.CreateUser(c, request)
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
