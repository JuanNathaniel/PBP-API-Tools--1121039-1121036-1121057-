package main

import (
	"fmt"
	"log"
	"net/http"

	"API_Exploration_3/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// controllers.RunScheduler()
	controllers.Connect()
	//controllers.sendEmail()

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
