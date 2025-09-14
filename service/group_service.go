package service

import (
	"jima/config"
	api_entity "jima/entity/api"
	"jima/entity/model"
	"jima/helper"
	"jima/repository"

	"github.com/gin-gonic/gin"
)

type GroupService interface {
	CreateGroup(c *gin.Context, request api_entity.GroupCreateGroupRequest) (response *api_entity.GroupCreateGroupResponse, err error)
}

type groupService struct {
	config              config.Config
	userRepository      repository.UserRepository
	groupRepository     repository.GroupRepository
	userGroupRepository repository.UserGroupRepository
}

func NewGroupService(
	config config.Config,
	userRepo repository.UserRepository,
	groupRepo repository.GroupRepository,
	userGroupRepo repository.UserGroupRepository,
) GroupService {
	return &groupService{
		config:              config,
		userRepository:      userRepo,
		groupRepository:     groupRepo,
		userGroupRepository: userGroupRepo,
	}
}

func (s *groupService) CreateGroup(c *gin.Context, request api_entity.GroupCreateGroupRequest) (response *api_entity.GroupCreateGroupResponse, err error) {
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

	return &api_entity.GroupCreateGroupResponse{
		Serial: group.Serial,
		Name:   group.Name,
	}, nil
}
