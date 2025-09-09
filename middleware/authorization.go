package middleware

import (
	"jima/config"
	"jima/entity/model"
	"jima/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HeaderAuthorization = "Authorization"

	ContextUserAuth = "userAuth"
)

func Authorization(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(HeaderAuthorization)

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": helper.ErrUnauthorizedToken.Error(),
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

		c.Set(ContextUserAuth, claims)
		c.Next()
	}
}

func ValidateUserRole(allowedRoles ...model.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAuth, exists := c.Get(ContextUserAuth)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": helper.ErrUnauthorizedToken.Error(),
			})
			return
		}

		claims := userAuth.(*helper.Claims)

		for _, v := range allowedRoles {
			if claims.Role == string(v) {
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
