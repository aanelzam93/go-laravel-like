package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-laravel-like/app/models"
	"go-laravel-like/app/helpers"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.ResponseJSON(c, 401, "Unauthorized", nil)
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		userID, err := models.VerifyToken(tokenString)
		if err != nil {
			helpers.ResponseJSON(c, 401, "Unauthorized", nil)
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}