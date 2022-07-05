package handlers

import (
	"comment-service/model"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	var commentCreationDTO model.CommentCreationDTO
	_ = json.NewDecoder(request.Body).Decode(&commentCreationDTO)

	userId := getUserId(request)

	ctx, cancel := context.WithTimeout(request.Context(), time.Second*20)
	defer cancel()

	//nema roditelja, root komentar
	if commentCreationDTO.ParentCommentId == "" {
		res, err := mongoHandler.CommentCollection.InsertOne(ctx, bson.D{
			{"user", commentCreationDTO.User},
			{"userId", userId},
			{"content", commentCreationDTO.Content},
			{"drinkId", commentCreationDTO.DrinkId},
			{"children", []model.CommentSavedDTO{}},
		})
		if err != nil {
			writeInternalServerError(writer, request, err)
		} else {
			json.NewEncoder(writer).Encode(model.ToCommentDTOFromCommentCreationDTO(&commentCreationDTO, getDocumentId(res)))
		}
	} else {
		objectId, _ := primitive.ObjectIDFromHex(commentCreationDTO.ParentCommentId)
		filter := bson.M{
			"$or": []bson.M{{"_id": objectId}, {"id": objectId}}}
		//options := options.FindOptions().SetProjection(bson.D{{"type", 1}, {"rating", 1}, {"_id", 0}})
		findOneOptions := &options.FindOneOptions{}
		findOneOptions.SetProjection(bson.D{
			{"id", 1},
			{"_id", 1},
			{"children", 1},
			{"user", 1},
			{"content", 1},
			{"userId", 1},
			{"drinkId", 1},
		})
		//opts := options.FindOneOptions().SetProjection(bson.D{{"type", 1}, {"rating", 1}, {"_id", 0}})
		result := mongoHandler.CommentCollection.FindOne(ctx, filter, findOneOptions)
		parentComment := model.CommentSavedDTO{}
		result.Decode(parentComment)
		fmt.Printf("parent comment id: %s\n", parentComment.Id.Hex())
		fmt.Println("parent comment content: " + parentComment.Content)
		fmt.Printf("parent comment children num: %d\n", len(parentComment.Children))
		newComment := model.ToCommentSavedDTOFromCommentCreationDTO(commentCreationDTO)
		newCommentId := primitive.NewObjectID()
		newComment.Id = newCommentId
		newComment.UserID = uint(userId)
		parentComment.Children = append(parentComment.Children, newComment)

		update := bson.D{{"$set", bson.D{{"children", parentComment.Children}}}}
		_, err3 := mongoHandler.CommentCollection.UpdateOne(ctx, filter, update)
		if err3 != nil {
			writeInternalServerError(writer, request, err3)
		} else {
			json.NewEncoder(writer).Encode(model.ToCommentDTOFromCommentSavedDTO(newComment))
		}
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
