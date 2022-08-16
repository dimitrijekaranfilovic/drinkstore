package main

import (
	"comment-service/handlers"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error creating mongo client: %+v", err)
	}
	defer client.Disconnect(ctx)
	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %+v", err)
	}

	handler := handlers.MongoHandler{
		CommentCollection: client.Database("ntp-comment-service").Collection("comments"),
		ReportCollection:  client.Database("ntp-comment-service").Collection("reports"),
	}

	port := "127.0.0.1:8083"
	router := mux.NewRouter()

	//user
	router.HandleFunc("/api/comments", handler.CreateComment).Methods("POST") 
	//user
	router.HandleFunc("/api/comments/reports/create", handler.ReportComment).Methods("POST") 
	//svi
	router.HandleFunc("/api/comments/for-drink/{drinkId}", handler.GetCommentsForDrink).Methods("GET") 
	//admin
	router.HandleFunc("/api/comments/reports/get", handler.GetAllReports).Methods("GET") 
	//admin
	router.HandleFunc("/api/comments/{id}", handler.DeleteComment).Methods("DELETE") 
	//admin
	router.HandleFunc("/api/comments/reports/{id}", handler.DeleteReport).Methods("DELETE")

	fmt.Println("Listening on: " + port)
	log.Fatalln(http.ListenAndServe(port, router))
}
