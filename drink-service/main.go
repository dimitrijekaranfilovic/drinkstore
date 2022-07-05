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

	//TODO: dodati dobavljanje slika pica
	//svi
	router.HandleFunc("/api/drinks", requestHandler.GetDrinks).Methods("GET")
	//svi
	router.HandleFunc("/api/drinks/{id}", requestHandler.GetSingleDrink).Methods("GET")
	//admin
	router.HandleFunc("/api/drinks", requestHandler.CreateDrink).Methods("POST")
	//admin
	router.HandleFunc("/api/drinks/{id}/images", requestHandler.CreateUpdateDrinkImage).Methods("POST")
	//admin
	router.HandleFunc("/api/drinks/{id}", requestHandler.UpdateDrink).Methods("PUT")
	//admin
	router.HandleFunc("/api/drinks/{id}", requestHandler.DeleteDrink).Methods("DELETE")
	//user
	router.HandleFunc("/api/drinks/{id}/grades", requestHandler.CreateUserGrade).Methods("POST")
	//user
	router.HandleFunc("/api/drinks/{drinkId}/grades/{gradeId}", requestHandler.UpdateUserGrade).Methods("PUT")
	//user
	router.HandleFunc("/api/drinks/{drinkId}/grades/{gradeId}", requestHandler.DeleteUserGrade).Methods("DELETE")
	//user
	router.HandleFunc("/api/drinks/{id}/grade-for-user", requestHandler.CheckDrinkGradeForUser).Methods("GET")
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:9090"},
		AllowedMethods: []string{"OPTIONS", "HEAD", "GET", "POST", "PUT", "DELETE"},
	})

	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, corsHandler.Handler(router)))

}
