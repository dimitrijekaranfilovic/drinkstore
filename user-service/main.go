package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/database"
	requestHandler "user-service/handlers"

	"github.com/gorilla/mux"
)

//TODO: provjeri komunikaciju sa ostalim servisima

func main() {
	database.ConnectToDatabase()

	port := ":8081"

	router := mux.NewRouter()

	//svi
	router.HandleFunc("/api/users/register", requestHandler.RegisterUser).Methods("POST") 
	//svi
	router.HandleFunc("/api/users/authenticate", requestHandler.Authenticate).Methods("POST") 
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
