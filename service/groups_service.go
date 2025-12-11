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
	GetGroupBySerial(c *gin.Context, request api_entity.GroupsGetGroupBySerialRequest) (response *api_entity.GroupsGetGroupBySerialResponse, err error)
	GetGroupsByUserSerial(c *gin.Context, request api_entity.GroupsGetGroupsByUserSerialRequest) (response *api_entity.GroupsGetGroupsByUserSerialResponse, err error)
	GetGroupMembers(c *gin.Context, request api_entity.GroupsGetGroupMembersRequest) (response *api_entity.GroupsGetGroupMembersResponse, err error)
	UpdateGroup(c *gin.Context, request api_entity.GroupsUpdateGroupRequest) (response *api_entity.GroupsUpdateGroupResponse, err error)
	UpdateGroupMemberRole(c *gin.Context, request api_entity.GroupsUpdateGroupMemberRoleRequest) (response *api_entity.GroupsUpdateGroupMemberRoleResponse, err error)
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

func (s *groupsService) GetGroupBySerial(c *gin.Context, request api_entity.GroupsGetGroupBySerialRequest) (response *api_entity.GroupsGetGroupBySerialResponse, err error) {
	// Check user a member of the group
	userGroup, err := s.userGroupRepository.GetUserGroup(request.UserAuthSerial, request.GroupSerial)
	if err != nil || userGroup == nil {
		return response, helper.ErrForbiddenUserAction
	}

	group, err := s.groupRepository.GetGroupBySerial(request.GroupSerial)
	if err != nil {
		return response, err
	}

	return &api_entity.GroupsGetGroupBySerialResponse{
		Group: api_entity.Group{
			GroupSerial: group.Serial,
			Name:        group.Name,
			CreatedAt:   group.CreatedAt.String(),
			UpdatedAt:   group.UpdatedAt.String(),
		},
	}, nil
}

func (s *groupsService) GetGroupsByUserSerial(c *gin.Context, request api_entity.GroupsGetGroupsByUserSerialRequest) (response *api_entity.GroupsGetGroupsByUserSerialResponse, err error) {
	userGroups, err := s.userGroupRepository.GetUserGroups(request.UserSerial)
	if err != nil {
		return response, err
	}

	response = &api_entity.GroupsGetGroupsByUserSerialResponse{}
	response.Groups = make([]api_entity.Group, len(userGroups))

	for i, userGroup := range userGroups {
		response.Groups[i] = api_entity.Group{
			GroupSerial: userGroup.Group.Serial,
			Name:        userGroup.Group.Name,
			CreatedAt:   userGroup.Group.CreatedAt.String(),
			UpdatedAt:   userGroup.Group.UpdatedAt.String(),
		}
	}

	return response, nil
}

func (s *groupsService) GetGroupMembers(c *gin.Context, request api_entity.GroupsGetGroupMembersRequest) (response *api_entity.GroupsGetGroupMembersResponse, err error) {
	// Check user a member of the group
	userGroup, err := s.userGroupRepository.GetUserGroup(request.UserAuthSerial, request.GroupSerial)
	if err != nil || userGroup == nil {
		return response, helper.ErrForbiddenUserAction
	}

	userGroups, err := s.userGroupRepository.GetUserGroupMembersByGroupSerial(request.GroupSerial)
	if err != nil {
		return response, err
	}

	response = &api_entity.GroupsGetGroupMembersResponse{}
	response.GroupMembers = make([]api_entity.GroupMember, len(userGroups))

	for i, userGroup := range userGroups {
		response.GroupMembers[i] = api_entity.GroupMember{
			UserGroupSerial: userGroup.Serial,
			UserSerial:      userGroup.UserSerial,
			UserName:        userGroup.User.Name,
			Role:            string(userGroup.Role),
			MemberSince:     userGroup.CreatedAt.String(),
		}
	}

	return response, nil
}

func (s *groupsService) UpdateGroup(c *gin.Context, request api_entity.GroupsUpdateGroupRequest) (response *api_entity.GroupsUpdateGroupResponse, err error) {
	// Check user a member of the group
	userGroup, err := s.userGroupRepository.GetUserGroup(request.UserAuthSerial, request.GroupSerial)
	if err != nil || userGroup == nil || userGroup.Role != model.UserGroupRoleManager {
		return response, helper.ErrForbiddenUserAction
	}

	updatedGroup, err := s.groupRepository.UpdateGroup(request)
	if err != nil {
		return response, err
	}

	return &api_entity.GroupsUpdateGroupResponse{
		Group: api_entity.Group{
			GroupSerial: updatedGroup.Serial,
			Name:        updatedGroup.Name,
			CreatedAt:   updatedGroup.CreatedAt.String(),
			UpdatedAt:   updatedGroup.UpdatedAt.String(),
		},
	}, nil
}

func (s *groupsService) UpdateGroupMemberRole(c *gin.Context, request api_entity.GroupsUpdateGroupMemberRoleRequest) (response *api_entity.GroupsUpdateGroupMemberRoleResponse, err error) {
	authUserGroup, err := s.userGroupRepository.GetUserGroup(request.UserAuthSerial, request.GroupSerial)
	if err != nil || authUserGroup == nil || authUserGroup.Role != model.UserGroupRoleManager {
		return response, helper.ErrForbiddenUserAction
	}
	if request.UserAuthSerial == request.UserSerial && request.Role != string(model.UserGroupRoleManager) {
		// check if current user is the only manager in the group
		managers, err := s.userGroupRepository.GetManagersInGroup(request.GroupSerial)
		if err != nil {
			return response, err
		}
		if len(managers) == 1 && managers[0].UserSerial == request.UserAuthSerial {
			return response, helper.ErrForbiddenUserAction
		}
	}

	// check is updated user a member of group
	userGroup, err := s.userGroupRepository.GetUserGroup(request.UserSerial, request.GroupSerial)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response, err
	}
	if userGroup == nil {
		return nil, helper.ErrUserNotInGroup
	}

	err = s.userGroupRepository.UpdateUserGroupRole(request.GroupSerial, request.UserSerial, request.Role)
	if err != nil {
		return response, err
	}

	// check is updated user a member of group
	updatedUserGroup, err := s.userGroupRepository.GetUserGroup(request.UserSerial, request.GroupSerial)
	if err != nil {
		return response, err
	}

	return &api_entity.GroupsUpdateGroupMemberRoleResponse{
		GroupMember: api_entity.GroupMember{
			UserGroupSerial: updatedUserGroup.Serial,
			UserSerial:      updatedUserGroup.UserSerial,
			UserName:        updatedUserGroup.User.Name,
			Role:            string(updatedUserGroup.Role),
			MemberSince:     updatedUserGroup.CreatedBy,
		},
	}, nil
}
