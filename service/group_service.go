package service

import (
	"fmt"
	"jima/config"
	api_entity "jima/entity/api"
	"jima/entity/model"
	"jima/helper"
	"jima/repository"
)

type GroupService interface {
	CreateGroup(request api_entity.GroupCreateGroupRequest) (response *api_entity.GroupCreateGroupResponse, err error)
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

func (s *groupService) CreateGroup(request api_entity.GroupCreateGroupRequest) (response *api_entity.GroupCreateGroupResponse, err error) {
	// Create group
	group := &model.Group{
		Serial: fmt.Sprintf("%s-%s", model.GroupSerialPrefix, helper.GenerateSerialFromString(request.Name)),
		Name:   request.Name,
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
