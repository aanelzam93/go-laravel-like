package docs

import (
	"github.com/gin-gonic/gin"
)

// @title Go Laravel-Like API
// @version 1.0
// @description API Documentation for Go Laravel Like Project
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email aan.elzam93@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /

// Register godoc
// @Summary Register new user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body object true "Register user input"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/register [post]
func Register(c *gin.Context) {}

// Login godoc
// @Summary Login user
// @Description Login and retrieve JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body object true "Login user input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/login [post]
func Login(c *gin.Context) {}

// Profile godoc
// @Summary Get user profile
// @Description Get authenticated user profile
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/profile [get]
func Profile(c *gin.Context) {}