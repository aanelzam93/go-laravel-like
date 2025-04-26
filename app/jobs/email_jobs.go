package jobs

import (
	"log"
	"time"
)

func SendWelcomeEmail(email string) {
	time.Sleep(2 * time.Second)
	log.Printf("Welcome email sent to %s!\n", email)
}