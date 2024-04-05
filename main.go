package main

import (
	"fmt"
	"log"
	"net/http"

	//"API_Exploration_3/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Panggil sendEmail()
	if err := controllers.sendEmail(); err != nil {
		log.Fatal("Failed to send email:", err)
	}

	//controllers.RunScheduler()
	//controllers.Connect()

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
