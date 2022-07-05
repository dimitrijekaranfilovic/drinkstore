package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CommentCreationDTO struct {
	User            string `json:"user"`
	Content         string `json:"content"`
	ParentCommentId string `json:"parentCommentId"`
	DrinkId         uint   `json:"drinkId"`
}

type CommentSavedDTO struct {
	Id       primitive.ObjectID `json:"_id"`
	User     string             `json:"user"`
	Content  string             `json:"content"`
	DrinkId  uint               `json:"drinkId"`
	UserID   uint               `json:"userId"`
	Children []CommentSavedDTO  `json:"children"`
}

type CommentDTO struct {
	Id       string       `json:"id"`
	User     string       `json:"user"`
	Content  string       `json:"content"`
	DrinkId  uint         `json:"drinkId"`
	Children []CommentDTO `json:"children"`
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

func ToCommentDTOFromCommentCreationDTO(dto *CommentCreationDTO, id string) CommentDTO {
	return CommentDTO{
		Id:       id,
		User:     dto.User,
		Content:  dto.Content,
		DrinkId:  dto.DrinkId,
		Children: []CommentDTO{},
	}
}

func ToCommentSavedDTOFromCommentCreationDTO(dto CommentCreationDTO) CommentSavedDTO {
	return CommentSavedDTO{
		User:     dto.User,
		Content:  dto.Content,
		DrinkId:  dto.DrinkId,
		Children: []CommentSavedDTO{},
	}
}

func ToCommentDTOArrayFromCommentSavedDTOArray(dtos []CommentSavedDTO) []CommentDTO {
	if len(dtos) == 0 {
		return []CommentDTO{}
	} else {
		var items []CommentDTO
		for _, item := range dtos {
			itemId := item.Id.Hex()
			newItem := CommentDTO{
				Id:      itemId,
				User:    item.User,
				Content: item.Content,
				DrinkId: item.DrinkId,
			}
			newItem.Children = ToCommentDTOArrayFromCommentSavedDTOArray(item.Children)
			items = append(items, newItem)
		}
		return items
	}
}

func ToCommentDTOFromCommentSavedDTO(dto CommentSavedDTO) CommentDTO {

	return CommentDTO{
		Id:       dto.Id.Hex(),
		User:     dto.User,
		Content:  dto.Content,
		DrinkId:  dto.DrinkId,
		Children: ToCommentDTOArrayFromCommentSavedDTOArray(dto.Children),
	}
}
