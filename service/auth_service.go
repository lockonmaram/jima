package service

import (
	"errors"
	"jima/config"
	api_entity "jima/entity/api"
	"jima/entity/model"
	"jima/helper"
	"jima/repository"

	"gorm.io/gorm"
)

type AuthService interface {
	Authenticate(username, password string) (response *api_entity.AuthAuthenticateResponse, err error)
	Register(request api_entity.AuthRegisterRequest) (response *api_entity.AuthRegisterResponse, err error)
}

type authService struct {
	config         config.Config
	userRepository repository.UserRepository
}

func NewAuthService(
	config config.Config,
	userRepo repository.UserRepository,
) AuthService {
	return &authService{
		config:         config,
		userRepository: userRepo,
	}
}

func (s *authService) Authenticate(userParam, password string) (response *api_entity.AuthAuthenticateResponse, err error) {
	// Get user by username or email
	user, err := s.userRepository.GetUserByUsernameOrEmail(userParam, userParam)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if user == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, helper.ErrUserNotFound
	}

	// Compare password
	if !helper.CompareHashAndPassword(user.Password, password) {
		return nil, helper.ErrInvalidPassword
	}

	// Generate token
	token, err := helper.GenerateJWT(s.config, user)
	if err != nil {
		return nil, err
	}

	return &api_entity.AuthAuthenticateResponse{
		Token:    token,
		Serial:   user.Serial,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (s *authService) Register(request api_entity.AuthRegisterRequest) (response *api_entity.AuthRegisterResponse, err error) {
	// Check if user already exists
	existingUser, err := s.userRepository.GetUserByUsernameOrEmail(request.Username, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, helper.ErrUserAlreadyExists
	}

	// Create user
	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}
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
