package helpers

import (
	"gopkg.in/gomail.v2"
	"os"
)

func SendEmail(to, subject, body string) error {
	if os.Getenv("MAIL_ENABLED") != "true" {
		return nil  
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	port := 587
	d := gomail.NewDialer(
		os.Getenv("MAIL_HOST"),
		port,
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
	)

	return d.DialAndSend(m)
}