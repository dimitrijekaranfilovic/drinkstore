package main

import (
	"comment-service/database"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	database.ConnectToDatabase()

	port := "127.0.0.1:8083"

	router := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:9090"},
		AllowedMethods:   []string{"OPTIONS", "HEAD", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, corsHandler.Handler(router)))
}
