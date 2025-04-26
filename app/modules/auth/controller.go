package auth

import (
	"github.com/gin-gonic/gin"
	"go-laravel-like/app/helpers"
)

type AuthController struct {
	Service *AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		Service: &AuthService{},
	}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if !helpers.BindAndValidate(c, &req) {
		return
	}

	if err := ctrl.Service.Register(req.Name, req.Email, req.Password); err != nil {
		helpers.ResponseJSON(c, 500, "Failed to register", err.Error())
		return
	}

	helpers.ResponseJSON(c, 201, "User registered successfully", nil)
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if !helpers.BindAndValidate(c, &req) {
		return
	}

	ok, err := ctrl.Service.Login(req.Email, req.Password)
	if !ok || err != nil {
		helpers.ResponseJSON(c, 401, "Invalid credentials", nil)
		return
	}

	helpers.ResponseJSON(c, 200, "Login successful", nil)
}