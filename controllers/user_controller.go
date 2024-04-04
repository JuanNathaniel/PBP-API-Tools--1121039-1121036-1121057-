package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// send email
func main() {
	d := gomail.NewDialer("joannanthaneil@gmail.com", 587, "Juan Nathaniel", "123456")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send emails using d.
	m := gomail.NewMessage()
	m.SetHeader("From", "joannanthaniel@gmail.com")
	m.SetHeader("To", "if-21039@students.ithb.ac.id")
	m.SetHeader("Subject", "Test email su")
	m.SetBody("text/plain", "Hello, this is a test email!")

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
