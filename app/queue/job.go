package queue

import (
	"encoding/json"
)

type Job struct {
	Name    string                 `json:"name"`
	Payload map[string]interface{} `json:"payload"`
	Delay   int                    `json:"delay"`  
	Retries int                    `json:"retries"`
}

func SerializeJob(job Job) ([]byte, error) {
	return json.Marshal(job)
}

func DeserializeJob(data []byte) (Job, error) {
	var job Job
	err := json.Unmarshal(data, &job)
	return job, err
}