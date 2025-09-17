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
	userController controller.UserController,
	groupController controller.GroupController,
) *gin.Engine {
	router := gin.Default()

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

	groupV1 := v1.Group("/group")
	groupV1.Use(middleware.Authorization(config))
	groupV1.POST("/", groupController.CreateGroup)

	userV1 := v1.Group("/user")
	userV1.Use(middleware.Authorization(config))
	userV1.POST("/", middleware.ValidateUserRole(model.UserRoleAdmin), userController.CreateUser)
	userV1.PUT("/:serial/profile", userController.UpdateUserProfile)
	userV1.PUT("/:serial/change-password", userController.ChangePassword)

	return router
}
