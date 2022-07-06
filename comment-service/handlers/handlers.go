package handlers

import (
	"comment-service/model"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type MongoHandler struct {
	CommentCollection *mongo.Collection
}

func getDocumentId(result *mongo.InsertOneResult) string {
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex()
	} else {
		return ""
	}
}

func (mongoHandler *MongoHandler) CreateComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(request.Context(), time.Second*20)
	defer cancel()

	var dto model.CommentCreationDTO
	_ = json.NewDecoder(request.Body).Decode(&dto)

	userId := getUserId(request)

	res, err := mongoHandler.CommentCollection.InsertOne(ctx, bson.D{
		{"user", dto.User},
		{"userId", userId},
		{"content", dto.Content},
		{"drinkId", dto.DrinkId},
	})
	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		json.NewEncoder(writer).Encode(model.ToCommentDTOFromCommentCreationDTO(&dto, getDocumentId(res)))
	}

}

//user id iz tokena
func (mongoHandler *MongoHandler) ReportComment(writer http.ResponseWriter, request *http.Request) {

}

func (mongoHandler *MongoHandler) GetCommentsForDrink(writer http.ResponseWriter, request *http.Request) {

}

func (mongoHandler *MongoHandler) GetAllReports(writer http.ResponseWriter, request *http.Request) {

}
func (mongoHandler *MongoHandler) DeleteComment(writer http.ResponseWriter, request *http.Request) {

}

func getUserId(request *http.Request) int {
	client := &http.Client{}
	newRequest, _ := http.NewRequest("GET", "http://127.0.0.1:8081/api/users/userId", nil)
	newRequest.Header.Add("Authorization", request.Header.Get("Authorization"))
	response, _ := client.Do(newRequest)

	var userIdDTO model.UserIdDTO
	_ = json.NewDecoder(response.Body).Decode(&userIdDTO)
	return userIdDTO.UserId
}
