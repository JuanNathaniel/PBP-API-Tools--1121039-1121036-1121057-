package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// send email
func main() {
	d := gomail.NewDialer("smtp.gmail.com", 587, "if-21039@students.ithb.ac.id", "ITHB2021")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send emails using d.
	m := gomail.NewMessage()
	m.SetHeader("From", "if-21039@students.ithb.ac.id")
	m.SetHeader("To", "joannanthaniel@gmail.com")
	m.SetHeader("Subject", "Test email su")
	m.SetBody("text/plain", "Hello, this is a test email!")

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
