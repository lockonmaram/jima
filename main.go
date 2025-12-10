package main

import (
	"fmt"
	"jima/config"
	"jima/controller"
	"jima/database"
	"jima/repository"
	"jima/router"
	"jima/service"
	"net/smtp"
)

func main() {
	// Init config & database
	config := config.Get()
	pgDB := database.NewPostgresDB(config)

	// Init smtp client
	smtpClient := initSMTPClient(config)

	// Repositories
	userRepository := repository.NewUserRepository(pgDB)
	groupRepository := repository.NewGroupRepository(pgDB)
	userGroupRepository := repository.NewUserGroupRepository(pgDB)

	// Client Services
	smtpService := service.NewSMTPService(smtpClient)

	// Services
	authService := service.NewAuthService(config, smtpService, userRepository)
	userService := service.NewUsersService(config, smtpService, userRepository)
	groupService := service.NewGroupsService(config, userRepository, groupRepository, userGroupRepository)

	// Controller
	authController := controller.NewAuthController(config, authService)
	userController := controller.NewUsersController(userService)
	groupController := controller.NewGroupsController(groupService)

	// Init Router
	r := router.InitRouter(
		config,
		authController,
		userController,
		groupController,
	)
	r.Run(fmt.Sprintf(":%d", config.Port))
}

func initSMTPClient(config config.Config) service.SMTPClient {
	return service.SMTPClient{
		Auth:    smtp.PlainAuth("", config.SMTPEmail, config.SMTPPassword, config.SMTPHost),
		Address: fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort),
		Email:   config.SMTPEmail,
		Name:    config.SMTPSenderName,
	}
}
