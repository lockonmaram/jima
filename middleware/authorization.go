package middleware

import (
	"jima/config"
	"jima/entity/model"
	"jima/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(helper.HeaderAuthorization)

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": helper.ErrTokenRequired.Error(),
			})
			return
		}

		claims, err := helper.ValidateJWT(config, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": helper.ErrUnauthorizedToken.Error(),
				"error":   err.Error(),
			})
			return
		}

		c.Set(helper.ContextUserAuth, claims)
		c.Next()
	}
}

func ValidateUserRole(allowedRoles ...model.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAuth := helper.GetUserAuthClaims(c)
		if userAuth == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": helper.ErrUnauthorizedToken.Error(),
			})
			return
		}

		for _, v := range allowedRoles {
			if userAuth.Role == string(v) {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": helper.ErrForbiddenUserAction.Error(),
		})
	}
}
