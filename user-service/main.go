package main

import (
	"fmt"
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
	router.HandleFunc("/api/users/authorize-admin", requestHandler.AuthorizeAdmin).Methods("GET")
	//samo ostali servisi
	router.HandleFunc("/api/users/authorize-user", requestHandler.AuthorizeUser).Methods("GET")
	//samo ostali servisi
	router.HandleFunc("/api/users/userId", requestHandler.GetUserIdFromJWT).Methods("GET")


	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, router))

}
