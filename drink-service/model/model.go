package model

type Drink struct {
	Id           uint    `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name" gorm:"unique"`
	ImagePath    string  `json:"imagePath"`
	Description  string  `json:"description"`
	Volume       float32 `json:"volume"`
	VolumeLabel  string  `json:"volumeLabel"`
	AverageGrade float32 `json:"averageGrade"`
}

type UserGrade struct {
	Id      uint  `json:"id" gorm:"primaryKey"`
	UserId  uint  `json:"userId"`
	Grade   uint8 `json:"grade"`
	DrinkId uint  `json:"drinkId"`
}

func ToDrinkFromDrinkCreationDTO(drinkDTO *DrinkCreationDTO) Drink {
	return Drink{
		Name:         drinkDTO.Name,
		Description:  drinkDTO.Description,
		Volume:       drinkDTO.Volume,
		VolumeLabel:  drinkDTO.VolumeLabel,
		AverageGrade: 0,
	}
}
