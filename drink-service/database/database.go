package database

import (
	"drink-service/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Driver *gorm.DB
var Err error

func ConnectToDatabase() {
	connectionString := "host=localhost user=postgres password=root dbname=ntp-drink-service port=5432 sslmode=disable"
	Driver, Err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if Err != nil {
		log.Fatalln("Failed to connect database")
	} else {
		fmt.Println("Connection successful.")
	}

	Err = Driver.Migrator().DropTable(&model.Drink{})
	if Err != nil {
		log.Fatalln("Failed to drop 'Drink' table")
	} else {
		fmt.Println("Dropped 'Drink' table.")
	}

	Err = Driver.Migrator().DropTable(&model.UserGrade{})

	if Err != nil {
		log.Fatalln("Failed to drop 'UserGrade' table.")
	} else {
		fmt.Println("Dropped 'UserGrade' table.")
	}

	Err = Driver.AutoMigrate(&model.Drink{}, &model.UserGrade{})
	if Err != nil {
		log.Fatalln("Migration failed.")
	} else {
		fmt.Println("Migration successful.")
	}

	drinks := []model.Drink{
		{
			Name:         "Vinjak",
			ImagePath:    "images/vinjak.jpg",
			Description:  "Samo za najjace",
			Volume:       0.75,
			VolumeLabel:  "l",
			AverageGrade: 4.8,
			Category:     "LIQUOR",
			Price:        1199,
		},
		{
			Name:         "Vranac",
			ImagePath:    "images/vranac.jpg",
			Description:  "Najbolje iz Crne Gore",
			Volume:       0.75,
			VolumeLabel:  "l",
			AverageGrade: 4.1,
			Category:     "WINE",
			Price:        799,
		},
		{
			Name:         "Zajecarsko",
			ImagePath:    "images/zajecarsko.JPG",
			Description:  "Zaboravi na Jelen",
			Volume:       0.5,
			VolumeLabel:  "l",
			AverageGrade: 5.0,
			Category:     "BEER",
			Price:        85,
		},
		{
			Name:         "Coca cola",
			ImagePath:    "images/coca_cola.jpg",
			Description:  "Klasika",
			Volume:       2,
			VolumeLabel:  "l",
			AverageGrade: 0.0,
			Category:     "CARBONATED",
			Price:        220,
		},
		{
			Name:         "Multivitamin",
			ImagePath:    "images/multivitamin.jpg",
			Description:  "Opis",
			Volume:       0.75,
			VolumeLabel:  "l",
			AverageGrade: 4.8,
			Category:     "NON_CARBONATED",
			Price:        3000,
		},
	}

	grades := []model.UserGrade{
		{
			UserId: 2, Grade: 5, DrinkId: 2,
		},
		{
			UserId: 2, Grade: 2, DrinkId: 1,
		},
	}

	for _, drink := range drinks {
		Driver.Create(&drink)
	}

	for _, grade := range grades {
		Driver.Create(&grade)
	}

	fmt.Println("Created data.")
}
