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

	port := ":8082"

	router := mux.NewRouter()
	//svi
	router.HandleFunc("/api/drinks/get", requestHandler.GetDrinks).Methods("GET")
	//svi
	router.HandleFunc("/api/drinks/get-unfiltered", requestHandler.GetUnfilteredDrinks).Methods("GET")
	//svi
	router.HandleFunc("/api/drinks/get/{id}", requestHandler.GetSingleDrink).Methods("GET") 
	//admin
	router.HandleFunc("/api/drinks/create", requestHandler.CreateDrink).Methods("POST") 
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
	//svi
	router.HandleFunc("/api/drinks/images/{imageName}", requestHandler.GetSingleImage).Methods("GET") 

	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, router))

}
