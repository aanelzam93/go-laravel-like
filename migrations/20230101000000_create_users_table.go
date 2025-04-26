package migrations

import (
	"gorm.io/gorm"
	"go-laravel-like/app/models"
)

func Migrate(db *gorm.DB) error {
	if db != nil {
		return db.AutoMigrate(&models.User{})
	}
	return nil
}