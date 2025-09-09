package controller

import (
	api_entity "jima/entity/api"
	"jima/helper"
	"jima/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupController interface {
	CreateGroup(c *gin.Context)
}

type groupController struct {
	groupService service.GroupService
}

func NewGroupController(groupService service.GroupService) GroupController {
	return &groupController{
		groupService,
	}
}

func (gc *groupController) CreateGroup(c *gin.Context) {
	request := api_entity.GroupCreateGroupRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	userAuth := helper.GetUserAuthClaims(c)
	request.UserSerial = userAuth.Serial

	response, err := gc.groupService.CreateGroup(request)
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
