package database

import (
	"comment-service/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Client *mongo.Client
var Err error

func ConnectToDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, Err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if Err = Client.Disconnect(ctx); Err != nil {
			panic(Err)
		}
	}()

	collection := Client.Database("ntp-comment-service").Collection("comments")
	comment := model.Comment{
		User:    "user",
		Content: "ne slazem se",
	}

	res, err := collection.InsertOne(ctx, bson.D{{"user", comment.User}, {"content", comment.Content}})
	id := res.InsertedID
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Inserted comment with id: %s\n", id)
	}
}
