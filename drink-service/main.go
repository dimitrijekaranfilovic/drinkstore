package main

import (
	"drink-service/database"
	requestHandler "drink-service/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	database.ConnectToDatabase()

	port := "127.0.0.1:8082"

	router := mux.NewRouter()

	//svi
	router.HandleFunc("/api/drinks", requestHandler.GetDrinks).Methods("GET") //povezano
	//svi
	router.HandleFunc("/api/drinks/{id}", requestHandler.GetSingleDrink).Methods("GET") //povezano
	//admin
	router.HandleFunc("/api/drinks", requestHandler.CreateDrink).Methods("POST") //povezano
	//admin
	router.HandleFunc("/api/drinks/{id}/images", requestHandler.CreateUpdateDrinkImage).Methods("POST") //povezano
	//admin
	router.HandleFunc("/api/drinks/{id}", requestHandler.UpdateDrink).Methods("PUT") //povezano
	//admin
	router.HandleFunc("/api/drinks/{id}", requestHandler.DeleteDrink).Methods("DELETE") //povezano
	//user
	router.HandleFunc("/api/drinks/{id}/grades", requestHandler.CreateUserGrade).Methods("POST")
	//user
	router.HandleFunc("/api/drinks/{drinkId}/grades/{gradeId}", requestHandler.UpdateUserGrade).Methods("PUT")
	//user
	router.HandleFunc("/api/drinks/{drinkId}/grades/{gradeId}", requestHandler.DeleteUserGrade).Methods("DELETE")
	//user
	router.HandleFunc("/api/drinks/{id}/grade-for-user", requestHandler.CheckDrinkGradeForUser).Methods("GET")
	//svi
	router.HandleFunc("/api/drinks/images/{imageName}", requestHandler.GetSingleImage).Methods("GET") //povezano
	corsHandler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:9090"},
		AllowedMethods:   []string{"OPTIONS", "HEAD", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, corsHandler.Handler(router)))

}
