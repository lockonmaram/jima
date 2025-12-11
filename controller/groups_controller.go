package controller

import (
	"errors"
	api_entity "jima/entity/api"
	"jima/helper"
	"jima/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupsController interface {
	CreateGroup(c *gin.Context)
	AddUserToGroup(c *gin.Context)
}

type groupsController struct {
	groupsService service.GroupsService
}

func NewGroupsController(groupsService service.GroupsService) GroupsController {
	return &groupsController{
		groupsService,
	}
}

func (gc *groupsController) CreateGroup(c *gin.Context) {
	request := api_entity.GroupsCreateGroupRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	userAuth := helper.GetUserAuthClaims(c)
	request.UserSerial = userAuth.Serial

	response, err := gc.groupsService.CreateGroup(c, request)
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

func (gc *groupsController) AddUserToGroup(c *gin.Context) {
	request := api_entity.GroupsAddUserToGroupRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	userAuth := helper.GetUserAuthClaims(c)
	request.UserAuthSerial = userAuth.Serial

	response, err := gc.groupsService.AddUserToGroup(c, request)
	if err != nil {
		if errors.Is(err, helper.ErrUserAlreadyInGroup) {
			helper.HandleResponse(c, helper.Response{
				Status: http.StatusBadRequest,
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
