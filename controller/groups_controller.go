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
	RemoveUserFromGroup(c *gin.Context)
	GetGroupDetail(c *gin.Context)
	GetUserGroups(c *gin.Context)
	GetGroupMembers(c *gin.Context)
	UpdateGroup(c *gin.Context)
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

func (gc *groupsController) RemoveUserFromGroup(c *gin.Context) {
	request := api_entity.GroupsRemoveUserFromGroupRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	userAuth := helper.GetUserAuthClaims(c)
	request.UserAuthSerial = userAuth.Serial

	response, err := gc.groupsService.RemoveUserFromGroup(c, request)
	if err != nil {
		if errors.Is(err, helper.ErrUserNotInGroup) {
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
		Status:  http.StatusOK,
		Message: response.Message,
	})
}

func (gc *groupsController) GetGroupDetail(c *gin.Context) {
	request := api_entity.GroupsGetGroupBySerialRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	userAuth := helper.GetUserAuthClaims(c)
	request.UserAuthSerial = userAuth.Serial

	response, err := gc.groupsService.GetGroupBySerial(c, request)
	if err != nil {
		helper.HandleResponse(c, helper.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	helper.HandleResponse(c, helper.Response{
		Status: http.StatusOK,
		Data:   response.Group,
	})
}

func (gc *groupsController) GetUserGroups(c *gin.Context) {
	request := api_entity.GroupsGetGroupsByUserSerialRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	if !helper.IsUserAdminOrSelf(c, request.UserSerial) {
		helper.HandleResponse(c, helper.Response{
			Status: http.StatusUnauthorized,
			Error:  helper.ErrForbiddenUserAction.Error(),
		})
		return
	}

	response, err := gc.groupsService.GetGroupsByUserSerial(c, request)
	if err != nil {
		helper.HandleResponse(c, helper.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	helper.HandleResponse(c, helper.Response{
		Status: http.StatusOK,
		Data:   response.Groups,
	})
}

func (gc *groupsController) GetGroupMembers(c *gin.Context) {
	request := api_entity.GroupsGetGroupMembersRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	userAuth := helper.GetUserAuthClaims(c)
	request.UserAuthSerial = userAuth.Serial

	response, err := gc.groupsService.GetGroupMembers(c, request)
	if err != nil {
		helper.HandleResponse(c, helper.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	helper.HandleResponse(c, helper.Response{
		Status: http.StatusOK,
		Data:   response.GroupMembers,
	})
}

func (gc *groupsController) UpdateGroup(c *gin.Context) {
	request := api_entity.GroupsUpdateGroupRequest{}
	if err := helper.HandleRequest(c, &request); err != nil {
		return
	}

	userAuth := helper.GetUserAuthClaims(c)
	request.UserAuthSerial = userAuth.Serial

	response, err := gc.groupsService.UpdateGroup(c, request)
	if err != nil {
		if errors.Is(err, helper.ErrInvalidRequest) {
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
		Data:   response.Group,
	})
}
