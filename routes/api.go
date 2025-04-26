package routes

import (
	"github.com/gin-gonic/gin"
	"go-laravel-like/app/controllers"
	"go-laravel-like/app/modules/auth"

	"go-laravel-like/app/middlewares"
)

func APIRoutes(router *gin.Engine) {
	userController := controllers.NewUserController()
	authController := auth.NewAuthController()


	api := router.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)

		// Protected routes
		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			auth.GET("/profile", userController.Profile)
		}
	}
}