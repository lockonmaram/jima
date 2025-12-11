package router

import (
	"jima/config"
	"jima/controller"
	"jima/entity/model"
	"jima/helper"
	"jima/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(
	config config.Config,
	authController controller.AuthController,
	usersController controller.UsersController,
	groupsController controller.GroupsController,
) *gin.Engine {
	router := gin.Default()

	// Load HTML Templates
	router.LoadHTMLGlob("view/**/*")

	// Health check
	router.GET("/health", func(c *gin.Context) {
		helper.HandleResponse(c, helper.Response{
			Message: "healthy",
		})
	})

	// API v1
	api := router.Group("/api")
	v1 := api.Group("/v1")

	authV1 := v1.Group("/auth")
	authV1.POST("/", authController.Authenticate)
	authV1.POST("/register", authController.Register)
	authV1.POST("/forgot-password", authController.ForgotPassword)
	authV1.POST("/set-password", authController.SetPassword)
	authV1.GET("/set-password", authController.SetPasswordPage)

	groupV1 := v1.Group("/groups")
	groupV1.Use(middleware.Authorization(config))
	groupV1.POST("/", groupsController.CreateGroup)
	groupV1.PUT("/:groupSerial/add-user/:userSerial", groupsController.AddUserToGroup)
	groupV1.DELETE("/:groupSerial/remove-user/:userSerial", groupsController.RemoveUserFromGroup)

	userV1 := v1.Group("/users")
	userV1.Use(middleware.Authorization(config))
	userV1.POST("/", middleware.ValidateUserRole(model.UserRoleAdmin), usersController.CreateUser)
	userV1.PUT("/:serial/profile", usersController.UpdateUserProfile)
	userV1.PUT("/:serial/change-password", usersController.ChangePassword)

	return router
}
