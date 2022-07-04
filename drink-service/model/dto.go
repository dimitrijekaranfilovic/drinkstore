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
	Volume      float32 `json:"volume"`
	VolumeLabel string  `json:"volumeLabel"`
}
