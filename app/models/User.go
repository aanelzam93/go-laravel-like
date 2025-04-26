package models

import (
	"time"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt"
	"go-laravel-like/config"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:100"`
	Email    string `gorm:"size:100;uniqueIndex"`
	Password string `gorm:"size:255"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  u.ID,
		"exp": time.Now().Add(config.JWT.Expire).Unix(),
	})

	return token.SignedString([]byte(config.JWT.Secret))
}

func VerifyToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT.Secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return uint(claims["id"].(float64)), nil
	}

	return 0, err
}