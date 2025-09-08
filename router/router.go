package router

import (
	"jima/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(
	authController controller.AuthController,
) *gin.Engine {
	router := gin.Default()

	// API v1
	api := router.Group("/api")
	v1 := api.Group("/v1")

	authV1 := v1.Group("/auth")
	authV1.POST("/", authController.Authenticate)
	authV1.POST("/register", authController.Register)

	return router
}
