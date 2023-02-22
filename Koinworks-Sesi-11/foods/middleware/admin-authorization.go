package middleware

import (
	"foods/utils/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)

		if userData["role"] != "ADMIN" {
			theErr := error_utils.NewNotAuthorized("you're not allowed to execute the action")
			c.AbortWithStatusJSON(theErr.Status(), theErr.Message())
			return
		}

		c.Next()
	}
}
