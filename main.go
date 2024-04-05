package main

import (
	"context"
	"log"
	"time"

	"API_Exploration_3/controllers"
	"API_Exploration_3/models"

	"github.com/go-co-op/gocron"

	_ "github.com/go-sql-driver/mysql"

	"crypto/tls"
	// "log"

	"gopkg.in/gomail.v2"
)

func main() {
	// router := mux.NewRouter()

	// Panggil sendEmail()
	// router.HandleFunc("/v2/insertuser2", controllers.sendEmail).Methods("GET")

	RunScheduler()
	// sendEmail()
	// getTodayNews()
	//controllers.Connect()

	// http.Handle("/", router)
	// fmt.Println("Connected to port 8080")
	// log.Println("Connected to port 8080")
	// log.Fatal(http.ListenAndServe(":8080", router))
}

// GoCron
func RunScheduler() {

	schedule := gocron.NewScheduler(time.UTC)

	schedule.StartAsync()

	// schedule.Every().Seconds().Do(func() {
	// 	log.Println("1")
	// 	//Go Routines
	// 	getTodayNews()
	// 	//Mock Redis Example
	// 	//go sendAdvertisement()
	// })
	sendAdvertisement()

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
	db := controllers.Connect()
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

// Go Redis
func sendAdvertisement() {
	var ctx = context.Background()
	rdb := controllers.ConnectRedis()
	db := controllers.Connect()
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
