package utils

import (
	"os"
	"log"
	"net/smtp"
	"server/models"
)

func SendEmail(email models.Email) {
	from := email.Sender
	to := []string{email.Receiver}
	msg := []byte(email.Body)

    user := os.Getenv("MAILTRAP_USER")
    password := os.Getenv("MAILTRAP_PASSWORD")
	
	port := "2525"
    host := "smtp.mailtrap.io"
	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully sent mail to all users")
}
