package queue

import (
	"time"
	"go-laravel-like/config"
	"log"
)

func DispatchJob(job Job) error {
	data, err := SerializeJob(job)
	if err != nil {
		return err
	}

	if job.Delay > 0 {
		delayKey := "delay:" + job.Name
		err = config.RedisClient.Set(config.RedisCtx, delayKey, data, time.Duration(job.Delay)*time.Second).Err()
		log.Println("Job delayed:", job.Name)
		return err
	}

	err = config.RedisClient.RPush(config.RedisCtx, "job_queue", data).Err()
	if err != nil {
		log.Println("Failed to dispatch job:", job.Name)
		return err
	}
	log.Println("Job dispatched:", job.Name)
	return nil
}