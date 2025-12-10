package service

import (
	"jima/config"
	api_entity "jima/entity/api"
	"jima/entity/model"
	"jima/helper"
	"jima/repository"

	"github.com/gin-gonic/gin"
)

type GroupsService interface {
	CreateGroup(c *gin.Context, request api_entity.GroupsCreateGroupRequest) (response *api_entity.GroupsCreateGroupResponse, err error)
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
