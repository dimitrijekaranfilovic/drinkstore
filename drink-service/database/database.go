package database

import (
	"drink-service/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
			Name: "Vinjak", ImagePath: "images/vinjak.jpg", Description: "Samo za najjace", Volume: 0.75, VolumeLabel: "l", AverageGrade: 0.0,
		},
		{
			Name: "Coca cola", ImagePath: "images/coca_cola.png", Description: "Klasika", Volume: 2, VolumeLabel: "l", AverageGrade: 0.0,
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
