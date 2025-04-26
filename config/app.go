package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name  string
	Env   string
	Debug bool
	Port  string
}

type JWTConfig struct {
	Secret string
	Expire time.Duration
}

var App AppConfig
var JWT JWTConfig

func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	App = AppConfig{
		Name:  GetEnv("APP_NAME", "GoLaravelLike"),
		Env:   GetEnv("APP_ENV", "development"),
		Debug: GetEnvAsBool("APP_DEBUG", true),
		Port:  GetEnv("APP_PORT", ":8000"),
	}

	expire, _ := strconv.Atoi(GetEnv("JWT_EXPIRE", "3600"))
	JWT = JWTConfig{
		Secret: GetEnv("JWT_SECRET", ""),
		Expire: time.Duration(expire) * time.Second,
	}
}


func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}


func GetEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "true"
	}
	return defaultValue
}