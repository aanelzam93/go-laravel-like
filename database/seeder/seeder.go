package seeder

import (
	"go-laravel-like/app/models"
	"go-laravel-like/config"
	"log"
)

func SeedUsers() {
	db := config.GetSQLDB()

	users := []models.User{
		{Name: "User One", Email: "user1@example.com"},
		{Name: "User Two", Email: "user2@example.com"},
		{Name: "User Three", Email: "user3@example.com"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Println("Error seeding user:", err)
		}
	}
	log.Println("User seeding complete!")
}