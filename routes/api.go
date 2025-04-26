package routes

import (
	"github.com/gin-gonic/gin"
	"go-laravel-like/app/controllers"
	"go-laravel-like/app/middlewares"
)

func APIRoutes(router *gin.Engine) {
	userController := controllers.NewUserController()

	api := router.Group("/api")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)

		// Protected routes
		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			auth.GET("/profile", userController.Profile)
		}
	}
}