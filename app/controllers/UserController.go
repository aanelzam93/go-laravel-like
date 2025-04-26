package controllers

import (
	"github.com/gin-gonic/gin"
	"go-laravel-like/app/models"
	"go-laravel-like/app/helpers"
	"go-laravel-like/config"
	"gorm.io/gorm"               
	"go.mongodb.org/mongo-driver/mongo"  
)

type UserController struct {
	SQLDB   *gorm.DB
	MongoDB *mongo.Client
}

func NewUserController() *UserController {
	return &UserController{
		SQLDB:    config.GetSQLDB(),
		MongoDB:  config.GetMongoDB(),
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseJSON(c, 400, "Invalid input", helpers.ValidationErrors(err))
		return
	}

	
	var existingUser models.User
	if err := uc.SQLDB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		helpers.ResponseJSON(c, 400, "Email already exists", nil)
		return
	}

	
	user := models.User{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := user.SetPassword(input.Password); err != nil {
		helpers.ResponseJSON(c, 500, "Error creating user", nil)
		return
	}

	if err := uc.SQLDB.Create(&user).Error; err != nil {
		helpers.ResponseJSON(c, 500, "Error creating user", nil)
		return
	}

	helpers.ResponseJSON(c, 201, "User registered successfully", nil)
}

func (uc *UserController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseJSON(c, 400, "Invalid input", helpers.ValidationErrors(err))
		return
	}

	
	var user models.User
	if err := uc.SQLDB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		helpers.ResponseJSON(c, 401, "Invalid credentials", nil)
		return
	}

	
	if !user.CheckPassword(input.Password) {
		helpers.ResponseJSON(c, 401, "Invalid credentials", nil)
		return
	}

	
	token, err := user.GenerateToken()
	if err != nil {
		helpers.ResponseJSON(c, 500, "Error generating token", nil)
		return
	}

	helpers.ResponseJSON(c, 200, "Login successful", gin.H{
		"token": token,
	})
}

func (uc *UserController) Profile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var user models.User
	if err := uc.SQLDB.First(&user, userID).Error; err != nil {
		helpers.ResponseJSON(c, 404, "User not found", nil)
		return
	}

	helpers.ResponseJSON(c, 200, "User profile", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}