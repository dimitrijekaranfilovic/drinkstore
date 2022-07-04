package model

//TODO: dodati kolicinu?
type Drink struct {
	Id           uint    `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name" gorm:"unique"`
	ImagePath    string  `json:"imagePath"`
	Description  string  `json:"description"`
	Volume       float32 `json:"volume"`
	VolumeLabel  string  `json:"volumeLabel"`
	AverageGrade float32 `json:"averageGrade"`
	Category     string  `json:"category"`
	Price        uint    `json:"price"`
}

type DrinkPage struct {
	Drinks        []Drink `json:"drinks"`
	Page          uint64  `json:"page"`
	SortCriteria  string  `json:"sortCriteria"`
	SortDirection string  `json:"sortDirection"`
	TotalPages    int64   `json:"totalPages"`
}

type UserGrade struct {
	Id      uint  `json:"id" gorm:"primaryKey"`
	UserId  uint  `json:"userId"`
	Grade   uint8 `json:"grade"`
	DrinkId uint  `json:"drinkId"`
}

func ToDrinkFromDrinkCreationDTO(drinkDTO *DrinkCreateUpdateDTO) Drink {
	return Drink{
		Name:         drinkDTO.Name,
		Description:  drinkDTO.Description,
		Volume:       drinkDTO.Volume,
		VolumeLabel:  drinkDTO.VolumeLabel,
		AverageGrade: 0,
		Category:     drinkDTO.Category,
		Price:        drinkDTO.Price,
	}
}
