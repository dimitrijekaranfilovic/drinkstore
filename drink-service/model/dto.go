package model

import "time"

//TODO: stavi u neki common
type Error struct {
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
	Path          string    `json:"path"`
	Status        int       `json:"status"`
	StatusMessage string    `json:"statusMessage"`
}

type DrinkCreateUpdateDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Volume      float32 `json:"volume,string"`
	VolumeLabel string  `json:"volumeLabel"`
	Category    string  `json:"category"`
	Price       uint    `json:"price,string"`
}

type UserGradeDTO struct {
	Grade uint8 `json:"grade"`
}

type UserCanGradeCheck struct {
	NumPurchases int64	`json:"num_purchases, string"`
}

//TODO: stavi u neki common
type UserIdDTO struct {
	UserId int `json:"userId"`
}

type GradeCheckDTO struct {
	GradeExists bool  `json:"gradeExists"`
	GradeValue  int16 `json:"gradeValue"`
	GradeId     uint  `json:"gradeId"`
}
