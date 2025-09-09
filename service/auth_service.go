package service

import (
	api_entity "jima/entity/api"
	"jima/entity/model"
	"jima/helper"
	"jima/repository"
)

type AuthService interface {
	// Authenticate(username, password string) (string, error)
	Register(request api_entity.AuthRegisterRequest) (response *api_entity.AuthRegisterResponse, err error)
}

type authService struct {
	userRepository      repository.UserRepository
	groupRepository     repository.GroupRepository
	userGroupRepository repository.UserGroupRepository
}

func NewUserService(
	userRepo repository.UserRepository,
	groupRepo repository.GroupRepository,
	userGroupRepo repository.UserGroupRepository,
) AuthService {
	return &authService{
		userRepository:      userRepo,
		groupRepository:     groupRepo,
		userGroupRepository: userGroupRepo,
	}
}

func (s *authService) Register(request api_entity.AuthRegisterRequest) (response *api_entity.AuthRegisterResponse, err error) {
	// Check if user already exists
	isExist, err := s.userRepository.CheckIsUserExist(request.Username, request.Email)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, helper.ErrUserAlreadyExists
	}

	// Create user
	// Hash password
	hashedPassword, err := helper.Hash(request.Password)
	if err != nil {
		return nil, err
	}
	// Generate serial from username
	serial := helper.GenerateSerialFromString(request.Username)

	user := model.User{
		Serial:   serial,
		Username: request.Username,
		Email:    request.Email,
		Name:     request.Name,
		Password: hashedPassword,
		Role:     request.Role,
	}

	err = s.userRepository.CreateUser(user)
	if err != nil {
		return nil, helper.ErrDatabase
	}

	return &api_entity.AuthRegisterResponse{
		Serial:   user.Serial,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Role:     user.Role,
	}, nil
}
