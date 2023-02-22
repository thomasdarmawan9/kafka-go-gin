package middleware

import (
	"foods/utils/error_utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secret_key = "SECRET_KEY"

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		theError := error_utils.NewNotAuthenticated("login to proceed")

		tokenStr := c.Request.Header.Get("Authorization")

		token, err := VerifyToken(tokenStr)

		if err != nil {
			c.AbortWithStatusJSON(theError.Status(), theError)
			return
		}

		if err != nil {
			c.AbortWithStatusJSON(theError.Status(), theError)
			return
		}

		c.Set("userData", token)
		c.Next()
	}
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error_utils.MessageErr) {
	if bearer := strings.HasPrefix(tokenStr, "Bearer"); !bearer {
		return nil, error_utils.NewNotAuthenticated("login to proceed")
	}

	stringToken := strings.Split(tokenStr, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, error_utils.NewNotAuthenticated("login to proceed")
		}
		return []byte(secret_key), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, error_utils.NewNotAuthenticated("login to proceed")
	}

	exp := int64(token.Claims.(jwt.MapClaims)["exp"].(float64))

	if (exp - time.Now().Unix()) <= 0 {
		return nil, error_utils.NewNotAuthenticated("login to proceed")
	}

	return token.Claims.(jwt.MapClaims), nil
}
