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

// // GoCron
// func RunScheduler() {
// 	schedule := gocron.NewScheduler(time.UTC)

// 	schedule.Every(1).Day().Do(func() {
// 		//Go Routines
// 		getTodayNews()
// 		//Mock Redis Example
// 		//go sendAdvertisement()
// 	})

// 	schedule.StartBlocking()
// 	//time.Sleep(10 * time.Second)
// }

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

// func getTodayNews() {
// 	db := Connect()
// 	defer db.Close()

// 	today := time.Now()

// 	query := `
// 			SELECT * FROM berita
// 			WHERE tanggal = ?`

// 	rows, err := db.Query(query, today.Format("2006-01-02"))
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	var berita Berita
// 	for rows.Next() {
// 		if err := rows.Scan(&berita.ID, &berita.Tanggal, &berita.Judul, &berita.Isi); err != nil {
// 			log.Println(err)
// 			return
// 		} else {

// 			query2 := `SELECT * FROM users`

// 			rows2, err := db.Query(query2)
// 			if err != nil {
// 				log.Println(err)
// 				return
// 			}

// 			var user User
// 			for rows2.Next() {
// 				if err := rows2.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
// 					log.Println(err)
// 					return
// 				} else {
// 					sendMail(user, berita)
// 				}
// 			}
// 		}
// 	}
// }

// // Go Redis
// func sendAdvertisement() {
// 	var ctx = context.Background()
// 	rdb := ConnectRedis()
// 	db := Connect()
// 	defer db.Close()

// 	promoCode, err := rdb.Get(ctx, "promoCode").Result()
// 	if err != nil {
// 		panic(err)
// 	}

// 	query := `SELECT * FROM users`

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	today := time.Now()
// 	berita := Berita{ID: 0, Tanggal: today.Format("2006-01-02"), Judul: "Todays Promo Code!", Isi: promoCode}
// 	var user User
// 	for rows.Next() {
// 		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
// 			log.Println(err)
// 			return
// 		} else {
// 			sendMail(user, berita)
// 		}
// 	}
// }
