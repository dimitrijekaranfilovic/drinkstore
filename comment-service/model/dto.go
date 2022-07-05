package model

import "time"

type CommentCreationDTO struct {
	User            string `json:"user"`
	Content         string `json:"content"`
	ParentCommentId string `json:"parentCommentId"`
	DrinkId         uint   `json:"drinkId"`
}

type CommentDTO struct {
	Id      string `json:"id"`
	User    string `json:"user"`
	Content string `json:"content"`
	DrinkId uint   `json:"drinkId"`
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

func ToCommentDTO(dto *CommentCreationDTO, id string) CommentDTO {
	return CommentDTO{
		Id:      id,
		User:    dto.User,
		Content: dto.Content,
		DrinkId: dto.DrinkId,
	}
}
