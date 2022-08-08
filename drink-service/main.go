package main

import (
	"drink-service/database"
	requestHandler "drink-service/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectToDatabase()

	port := "127.0.0.1:8082"

	router := mux.NewRouter()
	//svi
	router.HandleFunc("/api/drinks/get", requestHandler.GetDrinks).Methods("GET") //povezano
	//svi
	router.HandleFunc("/api/drinks/get-unfiltered", requestHandler.GetUnfilteredDrinks).Methods("GET")
	//svi
	router.HandleFunc("/api/drinks/get/{id}", requestHandler.GetSingleDrink).Methods("GET") //povezano
	//admin
	router.HandleFunc("/api/drinks/create", requestHandler.CreateDrink).Methods("POST") //povezano
	//admin
	router.HandleFunc("/api/drinks/{id}/images", requestHandler.CreateUpdateDrinkImage).Methods("POST") //povezano
	//admin
	router.HandleFunc("/api/drinks/{id}", requestHandler.UpdateDrink).Methods("PUT") //povezano
	//admin
	router.HandleFunc("/api/drinks/{id}", requestHandler.DeleteDrink).Methods("DELETE") //povezano
	//user
	router.HandleFunc("/api/drinks/{id}/grades", requestHandler.CreateUserGrade).Methods("POST") //povezano
	//user
	router.HandleFunc("/api/drinks/{drinkId}/grades/{gradeId}", requestHandler.UpdateUserGrade).Methods("PUT") //povezano
	//user
	router.HandleFunc("/api/drinks/{drinkId}/grades/{gradeId}", requestHandler.DeleteUserGrade).Methods("DELETE") //povezano
	//user
	router.HandleFunc("/api/drinks/{id}/grade-for-user", requestHandler.CheckDrinkGradeForUser).Methods("GET") //povezano
	//svi
	router.HandleFunc("/api/drinks/images/{imageName}", requestHandler.GetSingleImage).Methods("GET") //povezano

	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, router))

}
