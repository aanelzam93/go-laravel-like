package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

func ConnectRedis() {
	if GetEnv("REDIS_ENABLED", "false") != "true" {
		log.Println("Redis is disabled in configuration.")
		return
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_HOST", "localhost:6379"),
		Password: "", 
		DB:       0,  
	})

	_, err := RedisClient.Ping(RedisCtx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}