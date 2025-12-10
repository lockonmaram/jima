package controller

import (
	api_entity "jima/entity/api"
	"jima/helper"
	"jima/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupsController interface {
	CreateGroup(c *gin.Context)
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
