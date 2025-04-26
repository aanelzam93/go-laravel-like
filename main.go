package main

import (
	"log"
	"go-laravel-like/config"
	"go-laravel-like/routes"
	"go-laravel-like/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDatabaseConfig()

	config.ConnectDatabases()

	if sqlDB := config.GetSQLDB(); sqlDB != nil {
		if err := migrations.Migrate(sqlDB); err != nil {
			log.Fatal("Failed to migrate database:", err)
		}
	}

	// Initialize Gin
	router := gin.Default()

	// Setup routes
	routes.WebRoutes(router)
	routes.APIRoutes(router)  

	// Start server
	log.Printf("Server running on port %s", config.App.Port)
	router.Run(config.App.Port)
}