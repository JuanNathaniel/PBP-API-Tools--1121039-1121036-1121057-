package controllers

import (
	"context"
	"log"
	"time"

	"API_Exploration_3/models"

	"github.com/go-co-op/gocron"

	_ "github.com/go-sql-driver/mysql"

	"crypto/tls"
	// "log"

	"gopkg.in/gomail.v2"
)

// GoCron
func RunScheduler() {

	schedule := gocron.NewScheduler(time.UTC)

	schedule.StartAsync()

	schedule.Every(1).Day().Do(func() {
		log.Println("1")
		//Go Routines
		getTodayNews()
		//Mock Redis Example
		//go sendAdvertisement()
	})

	// getTodayNews()

	schedule.StartBlocking()
	//time.Sleep(10 * time.Second)
}

func sendEmail(user models.User, news models.Berita) error {
	d := gomail.NewDialer("smtp.gmail.com", 587, "if-21039@students.ithb.ac.id", "zgns qqbk vyer oodi")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send emails using d.
	m := gomail.NewMessage()
	m.SetHeader("From", "if-21039@students.ithb.ac.id")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", news.Title)
	m.SetBody("text/plain", news.Text)

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}

func getTodayNews() {
	db := Connect()
	defer db.Close()

	query := `SELECT * FROM berita WHERE tanggal = ?`

	rows, err := db.Query(query, time.Now().Format("2006-01-02"))

	if err != nil {
		log.Println(err)
		return
	}

	var berita models.Berita
	for rows.Next() {
		if err := rows.Scan(&berita.ID, &berita.Date, &berita.Title, &berita.Text); err != nil {
			log.Println(err)

			return
		} else {

			query2 := `SELECT * FROM users`

			rows2, err := db.Query(query2)
			if err != nil {
				log.Println(err)
				return
			}

			var user models.User
			for rows2.Next() {
				if err := rows2.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
					log.Println(err)
					return
				} else {
					// sendEmail(user, berita)
					sendEmail(user, berita)
				}
			}
		}
	}
}

// func getBeritaSekarang() {
// 	db := Connect()
// 	defer db.Close()

// 	query := `SELECT * FROM berita WHERE tanggal = ?`

// 	rows, err := db.Query(query, time.Now().Format("2006-01-02"))

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	var beritaa []models.Beritaa
// 	for rows.Next() {
// 		var berita models.Berita
// 		if err := rows.Scan(&berita.ID, &berita.Date, &berita.Title, &berita.Text); err != nil {
// 			log.Println(err)

// 			return
// 		}
// 		beritaa = append(beritaa, berita)

// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	var response = m.UserResponse{}
// 	response.Status = http.StatusOK
// 	response.Message = "Success Update the Address for >min_age and <max_age"
// 	//response.Data = users
// 	json.NewEncoder(w).Encode(beritaa)
// }

// Go Redis
func sendAdvertisement() {
	var ctx = context.Background()
	rdb := ConnectRedis()
	db := Connect()
	defer db.Close()

	promoCode, err := rdb.Get(ctx, "promoCode").Result()
	if err != nil {
		panic(err)
	}

	query := `SELECT * FROM users`

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	// today := time.Now()
	berita := models.Berita{ID: 0, Date: time.Now().Format("2006-01-02"), Title: "Todays Promo Code!", Text: promoCode}
	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
			log.Println(err)
			return
		} else {
			sendEmail(user, berita)
		}
	}
}
