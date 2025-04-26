package queue

import (
	"context"
	"log"
	"time"
	"go-laravel-like/config"
)

func StartWorker() {
	if config.RedisClient == nil {
		log.Println("Queue Worker not started: Redis is disabled.")
		return
	}

	log.Println("Advanced Queue Worker started...")

	for {
		data, err := config.RedisClient.BLPop(context.Background(), 0*time.Second, "job_queue").Result()
		if err != nil {
			log.Println("Error popping job:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		jobData := []byte(data[1])
		job, err := DeserializeJob(jobData)
		if err != nil {
			log.Println("Failed to deserialize job:", err)
			continue
		}

		log.Printf("Processing Job: %s\n", job.Name)

		success := executeJob(job)

		if !success && job.Retries > 0 {
			log.Printf("Job %s failed, retrying (%d retries left)\n", job.Name, job.Retries-1)
			job.Retries--
			DispatchJob(job)
		}
	}
}

func executeJob(job Job) bool {
	log.Println("Executing payload:", job.Payload)
	time.Sleep(2 * time.Second)  
	return true  
}