package routes

import "github.com/gin-gonic/gin"

func WebRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Laravel-like Framework with Gin!",
		})
	})
}