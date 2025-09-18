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

type UserService interface {
	CreateUser(c *gin.Context, request api_entity.UserCreateUserRequest) (response *api_entity.UserCreateUserResponse, err error)
	UpdateUserProfile(c *gin.Context, request api_entity.UserUpdateUserProfileRequest) (response *api_entity.UserUpdateUserProfileResponse, err error)
	ChangePassword(c *gin.Context, request api_entity.UserChangePasswordRequest) (err error)
}

type userService struct {
	config         config.Config
	smtpService    SMTPService
	userRepository repository.UserRepository
}

func NewUserService(
	config config.Config,
	smtpService SMTPService,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		config:         config,
		smtpService:    smtpService,
		userRepository: userRepo,
	}
}

func (s *userService) CreateUser(c *gin.Context, request api_entity.UserCreateUserRequest) (response *api_entity.UserCreateUserResponse, err error) {
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
	serial := helper.GenerateSerialFromString(model.UserSerialPrefix, request.Username)

	user := model.User{
		Serial:    serial,
		Username:  request.Username,
		Email:     request.Email,
		Name:      request.Name,
		Password:  hashedPassword,
		Role:      request.Role,
		CreatedBy: helper.GetUserAuthClaims(c).Serial,
	}

	err = s.userRepository.CreateUser(user)
	if err != nil {
		return nil, helper.ErrDatabase
	}

	go s.smtpService.SendMail(
		user.Email,
		string(helper.SMTP_SubjectRegisterSuccess),
		helper.GenerateSMTPTemplate(helper.SMTP_TemplateRegisterSuccess, user.Name),
	)

	return &api_entity.UserCreateUserResponse{
		Serial:   user.Serial,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Role:     user.Role,
	}, nil
}

func (s *userService) UpdateUserProfile(c *gin.Context, request api_entity.UserUpdateUserProfileRequest) (response *api_entity.UserUpdateUserProfileResponse, err error) {
	// Check update eligibility from auth
	if !helper.IsUserAdminOrSelf(c, request.Serial) {
		return nil, helper.ErrForbiddenUserAction
	}

	// Check if user already exists
	_, err = s.userRepository.GetUserBySerial(request.Serial)
	if err != nil {
		return nil, err
	}

	// Update user profile
	updatePayload := make(map[string]any)
	if request.Name != "" {
		updatePayload["name"] = request.Name
	}

	// Return error when no field update is valid
	if len(updatePayload) < 1 {
		return nil, helper.ErrInvalidRequest
	}

	user, err := s.userRepository.UpdateUserBySerial(request.Serial, updatePayload)
	if err != nil {
		return nil, helper.ErrDatabase
	}

	return &api_entity.UserUpdateUserProfileResponse{
		Name: user.Name,
	}, nil
}

func (s *userService) ChangePassword(c *gin.Context, request api_entity.UserChangePasswordRequest) (err error) {
	// Check update eligibility from auth
	if !helper.IsUserAdminOrSelf(c, request.Serial) {
		return helper.ErrForbiddenUserAction
	}

	// Check if user already exists
	user, err := s.userRepository.GetUserBySerial(request.Serial)
	if err != nil {
		return err
	}

	// Check if new & old password is the same
	if helper.CompareHashAndPassword(user.Password, request.Password) {
		return helper.ErrUnchangedPassword
	}

	// Hash new password
	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return err
	}

	err = s.userRepository.UpdateUserPasswordBySerialOrToken(request.Serial, hashedPassword)
	if err != nil {
		return helper.ErrDatabase
	}

	return nil
}
