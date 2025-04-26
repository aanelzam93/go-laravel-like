package queue

import (
	"context"
	"log"
	"time"
	"go-laravel-like/config"
)

func StartWorker() {
	if config.RedisClient == nil {
		log.Println("Worker not started: Redis not connected.")
		return
	}

	log.Println("Queue Worker started...")

	for {
		job, err := config.RedisClient.BLPop(context.Background(), 0*time.Second, "job_queue").Result()
		if err != nil {
			log.Println("Error popping job:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("Processing job:", job[1])
	}
}