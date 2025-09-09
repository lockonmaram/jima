package main

import (
	"fmt"
	"jima/config"
	"jima/controller"
	"jima/database"
	"jima/repository"
	"jima/router"
	"jima/service"
)

func main() {
	// Init config & database
	config := config.Get()
	pgDB := database.NewPostgresDB(config)

	// Repositories
	userRepository := repository.NewUserRepository(pgDB)
	groupRepository := repository.NewGroupRepository(pgDB)
	userGroupRepository := repository.NewUserGroupRepository(pgDB)

	// Services
	authService := service.NewAuthService(config, userRepository)
	groupService := service.NewGroupService(config, userRepository, groupRepository, userGroupRepository)

	// Controller
	authController := controller.NewAuthController(authService)
	groupController := controller.NewGroupController(groupService)

	// Init Router
	r := router.InitRouter(
		config,
		authController,
		groupController,
	)
	r.Run(fmt.Sprintf(":%d", config.Port))
}
