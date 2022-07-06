package handlers

import (
	"comment-service/model"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
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

	now := time.Now()
	res, err := mongoHandler.CommentCollection.InsertOne(ctx, bson.D{
		{"user", dto.User},
		{"userId", userId},
		{"content", dto.Content},
		{"drinkId", dto.DrinkId},
		{"postedAt", now},
	})
	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		json.NewEncoder(writer).Encode(model.ToCommentDTOFromCommentCreationDTO(&dto, getDocumentId(res), now))
	}

}
func (mongoHandler *MongoHandler) DeleteComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(request.Context(), time.Second*20)
	defer cancel()
	params := mux.Vars(request)
	commentId, _ := params["id"]
	objectId, _ := primitive.ObjectIDFromHex(commentId)
	_, err := mongoHandler.CommentCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}

}

func (mongoHandler *MongoHandler) GetCommentsForDrink(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(request.Context(), time.Second*20)
	defer cancel()
	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["drinkId"], 10, 64)
	opts := options.Find().SetSort(bson.D{{"postedAt", 1}})
	cursor, err := mongoHandler.CommentCollection.Find(ctx, bson.D{{"drinkId", drinkId}}, opts)
	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		var comments []model.CommentDTO
		for {
			if cursor.TryNext(ctx) {
				var comment model.CommentSaved
				cursor.Decode(&comment)
				newComment := model.CommentDTO{
					Id:       comment.Id.Hex(),
					User:     comment.User,
					Content:  comment.Content,
					DrinkId:  comment.DrinkId,
					PostedAt: comment.PostedAt,
				}
				comments = append(comments, newComment)
				continue
			}
			if cursor.ID() == 0 {
				break
			}
		}
		if comments == nil {
			json.NewEncoder(writer).Encode([]model.CommentDTO{})
		} else {
			json.NewEncoder(writer).Encode(comments)

		}
	}
}

//user id iz tokena
func (mongoHandler *MongoHandler) ReportComment(writer http.ResponseWriter, request *http.Request) {

}

func (mongoHandler *MongoHandler) GetAllReports(writer http.ResponseWriter, request *http.Request) {

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
