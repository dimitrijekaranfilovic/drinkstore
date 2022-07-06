package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"user-service/database"
	requestHandler "user-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectToDatabase()

	port := "127.0.0.1:8081"

	router := mux.NewRouter()

	//svi
	router.HandleFunc("/api/users/register", requestHandler.RegisterUser).Methods("POST") //povezano
	//svi
	router.HandleFunc("/api/users/authenticate", requestHandler.Authenticate).Methods("POST") //povezano
	//admin
	router.HandleFunc("/api/users/ban/{username}", requestHandler.BanUser).Methods("GET")
	//samo ostali servisi
	router.HandleFunc("/api/users/authorize", requestHandler.Authorize).Methods("GET")
	//samo ostali servisi
	router.HandleFunc("/api/users/userId", requestHandler.GetUserIdFromJWT).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:9090"},
		AllowedMethods:   []string{"OPTIONS", "HEAD", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, corsHandler.Handler(router)))

}
