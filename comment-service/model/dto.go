package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserCanGradeCheck struct {
	NumPurchases int64	`json:"num_purchases, string"`
}

type CommentCreationDTO struct {
	User    string `json:"user"`
	Content string `json:"content"`
	DrinkId uint   `json:"drinkId"`
}

type ReportCreationDTO struct {
	CommentContent string `json:"commentContent"`
	CommentId      string `json:"commentId"`
	PostedBy       string `json:"postedBy"`
	DrinkId        uint   `json:"drinkId,string"`
	ReportReason   string `json:"reportReason"`
	ReportedBy     string `json:"reportedBy"`
}

type ReportDTO struct {
	Id             string    `json:"id"`
	CommentContent string    `json:"commentContent"`
	CommentId      string    `json:"commentId"`
	PostedBy       string    `json:"postedBy"`
	DrinkId        uint      `json:"drinkId,string"`
	ReportReason   string    `json:"reportReason"`
	ReportedBy     string    `json:"reportedBy"`
	ReportedAt     time.Time `json:"reportedAt"`
}

type ReportSavedDTO struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	CommentContent string             `json:"commentContent"`
	CommentId      string             `json:"commentId"`
	PostedBy       string             `json:"postedBy"`
	DrinkId        uint               `json:"drinkId,string"`
	ReportReason   string             `json:"reportReason"`
	ReportedBy     string             `json:"reportedBy"`
	ReportedAt     time.Time          `json:"reportedAt"`
}

type CommentDTO struct {
	Id       string    `json:"id"`
	User     string    `json:"user"`
	Content  string    `json:"content"`
	DrinkId  uint      `json:"drinkId"`
	PostedAt time.Time `json:"postedAt"`
}

type CommentSaved struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	User     string             `json:"user"`
	Content  string             `json:"content"`
	DrinkId  uint               `json:"drinkId"`
	PostedAt time.Time          `json:"postedAt"`
}

type UserIdDTO struct {
	UserId int `json:"userId"`
}

type Error struct {
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
	Path          string    `json:"path"`
	Status        int       `json:"status"`
	StatusMessage string    `json:"statusMessage"`
}

func ToCommentDTOFromCommentCreationDTO(dto *CommentCreationDTO, id string, time time.Time) CommentDTO {
	return CommentDTO{
		Id:       id,
		User:     dto.User,
		Content:  dto.Content,
		DrinkId:  dto.DrinkId,
		PostedAt: time,
	}
}

func ToReportDTOFromReportCreationDTO(dto *ReportCreationDTO, id string, time time.Time) ReportDTO {
	return ReportDTO{
		Id:             id,
		CommentContent: dto.CommentContent,
		CommentId:      dto.CommentId,
		PostedBy:       dto.PostedBy,
		DrinkId:        dto.DrinkId,
		ReportReason:   dto.ReportReason,
		ReportedBy:     dto.ReportedBy,
		ReportedAt:     time,
	}
}
