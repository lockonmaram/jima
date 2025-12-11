package service

import (
	"errors"
	"jima/config"
	api_entity "jima/entity/api"
	"jima/entity/model"
	"jima/helper"
	"jima/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GroupsService interface {
	CreateGroup(c *gin.Context, request api_entity.GroupsCreateGroupRequest) (response *api_entity.GroupsCreateGroupResponse, err error)
	AddUserToGroup(c *gin.Context, request api_entity.GroupsAddUserToGroupRequest) (response *api_entity.GroupsAddUserToGroupResponse, err error)
	RemoveUserFromGroup(c *gin.Context, request api_entity.GroupsRemoveUserFromGroupRequest) (response *api_entity.GroupsRemoveUserFromGroupResponse, err error)
}

type groupsService struct {
	config              config.Config
	userRepository      repository.UserRepository
	groupRepository     repository.GroupRepository
	userGroupRepository repository.UserGroupRepository
}

func NewGroupsService(
	config config.Config,
	userRepo repository.UserRepository,
	groupRepo repository.GroupRepository,
	userGroupRepo repository.UserGroupRepository,
) GroupsService {
	return &groupsService{
		config:              config,
		userRepository:      userRepo,
		groupRepository:     groupRepo,
		userGroupRepository: userGroupRepo,
	}
}

func (s *groupsService) CreateGroup(c *gin.Context, request api_entity.GroupsCreateGroupRequest) (response *api_entity.GroupsCreateGroupResponse, err error) {
	// Create group
	group := &model.Group{
		Serial:    helper.GenerateSerialFromString(model.GroupSerialPrefix, request.Name),
		Name:      request.Name,
		CreatedBy: request.UserSerial,
	}

	group, err = s.groupRepository.CreateGroup(group, request.UserSerial)
	if err != nil {
		return nil, err
	}

	return &api_entity.GroupsCreateGroupResponse{
		Serial: group.Serial,
		Name:   group.Name,
	}, nil
}

func (s *groupsService) AddUserToGroup(c *gin.Context, request api_entity.GroupsAddUserToGroupRequest) (response *api_entity.GroupsAddUserToGroupResponse, err error) {
	// Check existing user group
	existingUserGroup, err := s.userGroupRepository.GetUserGroup(request.UserSerial, request.GroupSerial)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUserGroup != nil {
		return nil, helper.ErrUserAlreadyInGroup
	}

	// Check auth user status in group
	userGroup, err := s.userGroupRepository.GetUserGroup(request.UserAuthSerial, request.GroupSerial)
	if err != nil {
		return nil, err
	}
	if !helper.IsUserGroupManagerOrSelf(userGroup, request.UserSerial) {
		return nil, helper.ErrForbiddenUserAction
	}

	group, err := s.groupRepository.GetGroupBySerial(request.GroupSerial)
	if err != nil {
		return nil, err
	}

	createdUserGroup, err := s.userGroupRepository.AddUserToGroup(request.UserSerial, request.GroupSerial)
	if err != nil {
		return nil, err
	}

	return &api_entity.GroupsAddUserToGroupResponse{
		UserGroupSerial: createdUserGroup.Serial,
		GroupSerial:     createdUserGroup.GroupSerial,
		UserSerial:      createdUserGroup.UserSerial,
		GroupName:       group.Name,
	}, nil
}

func (s *groupsService) RemoveUserFromGroup(c *gin.Context, request api_entity.GroupsRemoveUserFromGroupRequest) (response *api_entity.GroupsRemoveUserFromGroupResponse, err error) {
	// Check existing user group
	existingUserGroup, err := s.userGroupRepository.GetUserGroup(request.UserSerial, request.GroupSerial)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response, err
	}
	if existingUserGroup == nil {
		return response, helper.ErrUserNotInGroup
	}

	// Check auth user status in group
	userGroup, err := s.userGroupRepository.GetUserGroup(request.UserAuthSerial, request.GroupSerial)
	if err != nil {
		return response, err
	}
	if !helper.IsUserGroupManagerOrSelf(userGroup, request.UserSerial) {
		return response, helper.ErrForbiddenUserAction
	}

	err = s.userGroupRepository.RemoveUserFromGroup(existingUserGroup.Serial)
	if err != nil {
		return response, err
	}

	return &api_entity.GroupsRemoveUserFromGroupResponse{
		Success: true,
		Message: helper.MsgUserHasBeenRemovedFromGroup,
	}, nil
}
