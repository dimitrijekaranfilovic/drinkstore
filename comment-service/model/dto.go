package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CommentCreationDTO struct {
	User    string `json:"user"`
	Content string `json:"content"`
	DrinkId uint   `json:"drinkId"`
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
