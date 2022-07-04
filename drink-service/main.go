package main

import (
	"drink-service/database"
	"drink-service/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database.ConnectToDatabase()

	port := "127.0.0.1:8082"

	router := mux.NewRouter()

	//TODO: dodati dobavljanje slika pica
	//svi
	router.HandleFunc("/api/drinks", handlers.GetDrinks).Methods("GET")
	//svi
	router.HandleFunc("/api/drinks/{id}", handlers.GetSingleDrink).Methods("GET")
	//admin
	router.HandleFunc("/api/drinks", handlers.CreateDrink).Methods("POST")
	//admin
	router.HandleFunc("/api/drinks/{id}/images", handlers.CreateUpdateDrinkImage).Methods("POST")
	//admin
	router.HandleFunc("/api/drinks/{id}", handlers.UpdateDrink).Methods("PUT")
	//admin
	router.HandleFunc("/api/drinks/{id}", handlers.DeleteDrink).Methods("DELETE")
	//user
	router.HandleFunc("/api/drinks/{id}/grades", handlers.CreateUserGrade).Methods("POST")
	//user
	router.HandleFunc("/api/drinks/{drinkId}/grades/{gradeId}", handlers.UpdateUserGrade).Methods("PUT")
	//user
	router.HandleFunc("/api/drinks/{drinkId}/grades/{gradeId}", handlers.DeleteUserGrade).Methods("DELETE")

	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, router))

}
