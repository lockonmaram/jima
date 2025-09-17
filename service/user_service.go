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
	CreateUser(c *gin.Context, request api_entity.UserCreateRequest) (response *api_entity.UserCreateResponse, err error)
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

func (s *userService) CreateUser(c *gin.Context, request api_entity.UserCreateRequest) (response *api_entity.UserCreateResponse, err error) {
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

	return &api_entity.UserCreateResponse{
		Serial:   user.Serial,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Role:     user.Role,
	}, nil
}
