package main
// @title Go Laravel-like API
// @version 1.0
// @description This is a sample server for Go Laravel-like Framework.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /api
// @schemes http
import (
	"log"
	"go-laravel-like/config"
	"go-laravel-like/routes"
	"go-laravel-like/migrations"
	"go-laravel-like/app/logs" 
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go-laravel-like/app/middlewares"
	_ "go-laravel-like/docs" 
		"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"go-laravel-like/app/queue"

)

func main() {
	config.LoadConfig()
	config.LoadDatabaseConfig()
	config.ConnectRedis()
	logs.InitLogger() 
	go queue.StartWorker()

	config.ConnectDatabases()

	if sqlDB := config.GetSQLDB(); sqlDB != nil {
		if err := migrations.Migrate(sqlDB); err != nil {
			logs.Error("Migration failed", zap.Error(err))
			log.Fatal("Failed to migrate database:", err)
		}
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.RecoveryWithLogger())

	routes.WebRoutes(router)
	routes.APIRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	

	logs.Info("Server running on port", zap.String("port", config.App.Port))
	router.Run(config.App.Port)
}