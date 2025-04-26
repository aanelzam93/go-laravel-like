package auth

import (
	"go-laravel-like/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) Register(name, email, password string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return config.GetSQLDB().Create(&user).Error
}

func (s *AuthService) Login(email, password string) (bool, error) {
	var user User
	db := config.GetSQLDB()
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return false, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil, err
}