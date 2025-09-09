package middleware

import (
	"jima/config"
	"jima/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HeaderAuthorization = "Authorization"
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

		c.Set("userAuth", claims)
		c.Next()
	}
}
