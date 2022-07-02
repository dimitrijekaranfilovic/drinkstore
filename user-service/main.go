package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user-service/database"
	"user-service/handlers"
)

func main() {
	database.ConnectToDatabase()

	port := ":8081"

	router := mux.NewRouter()

	router.HandleFunc("/api/users/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/users/authenticate", handlers.Authenticate).Methods("POST")
	router.HandleFunc("/api/users/ban/{id}", handlers.BanUser).Methods("GET") //vidi ovo malo bolje
	router.HandleFunc("/api/users/authorize", handlers.Authorize).Methods("GET")

	fmt.Println("Listening on port: " + port[1:])
	log.Fatalln(http.ListenAndServe(port, router))

}
