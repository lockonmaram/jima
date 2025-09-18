package service

import (
	"errors"
	"fmt"
	"jima/config"
	api_entity "jima/entity/api"
	"jima/entity/model"
	"jima/helper"
	"jima/repository"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthService interface {
	Authenticate(c *gin.Context, username, password string) (response *api_entity.AuthAuthenticateResponse, err error)
	Register(c *gin.Context, request api_entity.AuthRegisterRequest) (response *api_entity.AuthRegisterResponse, err error)
	SetPassword(c *gin.Context, request api_entity.AuthSetPasswordRequest) (err error)
	ForgotPassword(c *gin.Context, request api_entity.AuthForgotPasswordRequest) (err error)
}

type authService struct {
	config         config.Config
	smtpService    SMTPService
	userRepository repository.UserRepository
}

func NewAuthService(
	config config.Config,
	smtpService SMTPService,
	userRepo repository.UserRepository,
) AuthService {
	return &authService{
		config:         config,
		smtpService:    smtpService,
		userRepository: userRepo,
	}
}

func (s *authService) Authenticate(c *gin.Context, userParam, password string) (response *api_entity.AuthAuthenticateResponse, err error) {
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

func (s *authService) Register(c *gin.Context, request api_entity.AuthRegisterRequest) (response *api_entity.AuthRegisterResponse, err error) {
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
		Role:      string(model.UserRoleUser),
		CreatedBy: serial,
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

	return &api_entity.AuthRegisterResponse{
		Serial:   user.Serial,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
	}, nil
}

func (s *authService) SetPassword(c *gin.Context, request api_entity.AuthSetPasswordRequest) (err error) {
	// Check if user already exists
	user, err := s.userRepository.GetUserByPasswordToken(request.PasswordToken)
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

	err = s.userRepository.UpdateUserPasswordBySerialOrToken(request.PasswordToken, hashedPassword)
	if err != nil {
		return helper.ErrDatabase
	}

	return nil
}

func (s *authService) ForgotPassword(c *gin.Context, request api_entity.AuthForgotPasswordRequest) (err error) {
	// Check if user exists
	user, err := s.userRepository.GetUserByUsernameOrEmail(request.UserParam, request.UserParam)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Generate password token
	passwordToken, err := helper.HashPassword(fmt.Sprintf("%s:%v", user.Serial, time.Now().Unix()))
	if err != nil {
		return err
	}

	// Update user
	err = s.userRepository.SetPasswordToken(user.Serial, passwordToken)
	if err != nil {
		return err
	}

	// Send token by email
	resetPasswordURL := fmt.Sprintf("%s:%d/api/v1/auth/forgot-password?t=%s", s.config.BaseURL, s.config.Port, passwordToken)
	err = s.smtpService.SendMail(
		user.Email,
		string(helper.SMTP_SubjectRegisterSuccess),
		helper.GenerateSMTPTemplate(helper.SMTP_TemplateForgotPassword, resetPasswordURL),
	)
	if err != nil {
		return err
	}

	return nil
}
