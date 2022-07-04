package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/database"
	"user-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectToDatabase()

	port := "127.0.0.1:8081"

	router := mux.NewRouter()

	//svi
	router.HandleFunc("/api/users/register", handlers.RegisterUser).Methods("POST")
	//svi
	router.HandleFunc("/api/users/authenticate", handlers.Authenticate).Methods("POST")
	//admin
	router.HandleFunc("/api/users/ban/{id}", handlers.BanUser).Methods("GET")
	//samo ostali servisi
	router.HandleFunc("/api/users/authorize", handlers.Authorize).Methods("GET")
	//samo ostali servisi
	router.HandleFunc("/api/users/userId", handlers.GetUserIdFromJWT).Methods("GET")
	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, router))

}
