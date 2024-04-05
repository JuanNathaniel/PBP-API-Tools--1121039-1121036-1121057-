package controllers

import (
	// "log"
	// "time"

	// "context"

	// "github.com/go-co-op/gocron"
	"crypto/tls"
	"log"

	"gopkg.in/gomail.v2"
)

// send email
func sendEmail() error {
	d := gomail.NewDialer("smtp.gmail.com", 587, "if-21039@students.ithb.ac.id", "zgns qqbk vyer oodi")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send emails using d.
	m := gomail.NewMessage()
	m.SetHeader("From", "if-21039@students.ithb.ac.id")
	m.SetHeader("To", "wermichael211@gmail.com")
	m.SetHeader("Subject", "Test email ngehe")
	m.SetBody("text/plain", "Hello, this is a test email!")

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
