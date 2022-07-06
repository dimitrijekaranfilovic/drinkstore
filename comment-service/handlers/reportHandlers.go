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
	"time"
)

//user id iz tokena
func (mongoHandler *MongoHandler) ReportComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(request.Context(), time.Second*20)
	defer cancel()
	var dto model.ReportCreationDTO
	_ = json.NewDecoder(request.Body).Decode(&dto)

	now := time.Now()
	res, err := mongoHandler.ReportCollection.InsertOne(ctx, bson.D{
		{"commentContent", dto.CommentContent},
		{"commentId", dto.CommentId},
		{"postedBy", dto.PostedBy},
		{"drinkId", dto.DrinkId},
		{"reportReason", dto.ReportReason},
		{"reportedBy", dto.ReportedBy},
		{"reportedAt", now},
	})
	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		json.NewEncoder(writer).Encode(model.ToReportDTOFromReportCreationDTO(&dto, util.GetDocumentId(res), now))
	}

}

func (mongoHandler *MongoHandler) GetAllReports(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(request.Context(), time.Second*20)
	defer cancel()
	opts := options.Find().SetSort(bson.D{{"reportedAt", 1}})
	cursor, err := mongoHandler.ReportCollection.Find(ctx, bson.D{}, opts)
	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		var reports []model.ReportDTO
		for {
			if cursor.TryNext(ctx) {
				var report model.ReportSavedDTO
				cursor.Decode(&report)
				newReport := model.ReportDTO{
					Id:             report.Id.Hex(),
					CommentContent: report.CommentContent,
					CommentId:      report.CommentId,
					PostedBy:       report.PostedBy,
					DrinkId:        report.DrinkId,
					ReportReason:   report.ReportReason,
					ReportedBy:     report.ReportedBy,
					ReportedAt:     report.ReportedAt,
				}
				reports = append(reports, newReport)
				continue
			}
			if cursor.ID() == 0 {
				break
			}
		}
		if reports == nil {
			json.NewEncoder(writer).Encode([]model.ReportDTO{})
		} else {
			json.NewEncoder(writer).Encode(reports)

		}

	}
}

func (mongoHandler *MongoHandler) DeleteReport(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(request.Context(), time.Second*20)
	defer cancel()
	params := mux.Vars(request)
	reportId, _ := params["id"]
	objectId, _ := primitive.ObjectIDFromHex(reportId)
	_, err := mongoHandler.ReportCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}
