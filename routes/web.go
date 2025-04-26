
package routes

import (
	"github.com/gin-gonic/gin"
	"go-laravel-like/config"
)

func WebRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"db": config.GetSQLDB() != nil,
			"redis": config.RedisClient != nil,
		})
	})
}
