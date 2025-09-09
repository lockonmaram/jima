package router

import (
	"jima/config"
	"jima/controller"
	"jima/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(
	config config.Config,
	authController controller.AuthController,
) *gin.Engine {
	router := gin.Default()

	// API v1
	api := router.Group("/api")
	v1 := api.Group("/v1")

	authV1 := v1.Group("/auth")
	authV1.POST("/", authController.Authenticate)
	authV1.POST("/register", middleware.Authorization(config), authController.Register)

	return router
}
