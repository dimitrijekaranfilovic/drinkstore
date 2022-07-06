package handlers

import (
	"comment-service/model"
	"comment-service/util"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
)

func (mongoHandler *MongoHandler) CreateComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(request.Context(), time.Second*20)
	defer cancel()
	//TODO: pozovi purchase service
	var dto model.CommentCreationDTO
	_ = json.NewDecoder(request.Body).Decode(&dto)

	userId := util.GetUserId(request)

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
		json.NewEncoder(writer).Encode(model.ToCommentDTOFromCommentCreationDTO(&dto, util.GetDocumentId(res), now))
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
