package database

import (
	"fmt"
	"log"
	"os"
	"user-service/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Driver *gorm.DB
var Err error

func HashPassword(password string) string {
	bytePassword := []byte(password)
	hashedBytePassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Error hashing password.")
	}
	return string(hashedBytePassword)
}

func ConnectToDatabase() {
	databaseHost := os.Getenv("POSTGRES_HOST")
	if databaseHost == "" {
		databaseHost = "localhost"
	} 

	databaseUser := os.Getenv("POSTGRES_USER")
	if databaseUser == "" {
		databaseUser = "postgres"
	} 


	databasePassword := os.Getenv("POSTGRES_PASSWORD")
	if databasePassword == "" {
		databasePassword = "root"
	} 

	databaseName := os.Getenv("POSTGRES_DB")
	if databaseName == "" {
		databaseName = "ntp-user-service"
	} 

	

	//connectionString := "host=user-service-database user=postgres password=root dbname=ntp-user-service port=5432 sslmode=disable"
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", 
				databaseHost, 
				databaseUser, 
				databasePassword, 
				databaseName)
				
	Driver, Err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if Err != nil {
		log.Fatalln("Failed to connect database")
	} else {
		fmt.Println("Connection successful.")
	}

	Err = Driver.Migrator().DropTable(&model.User{})
	if Err != nil {
		log.Fatalln("Failed to drop 'User' table")
	} else {
		fmt.Println("Dropped 'User' table.")
	}

	Err = Driver.Migrator().DropTable(&model.Authority{})

	if Err != nil {
		log.Fatalln("Failed to drop 'Authority' table.")
	} else {
		fmt.Println("Dropped 'Authority' table.")
	}

	Err = Driver.AutoMigrate(&model.User{}, &model.Authority{})
	if Err != nil {
		log.Fatalln("Migration failed.")
	} else {
		fmt.Println("Migration successful.")
	}

	authorities := []model.Authority{
		{Name: "ADMIN"},
		{Name: "USER"},
	}

	users := []model.User{
		{
			Username: "admin", Password: HashPassword("123456"), FirstName: "Vidoje", LastName: "Gavrilovic", AuthorityId: 1,
			Banned: false,
			Address: model.Address{
				Street: "ulica", Number: "32", City: "Novi Sad", ZipCode: "21 000", Country: "Srbija",
			}},

		{
			Username: "user", Password: HashPassword("123456"), FirstName: "Jakov", LastName: "Matic", AuthorityId: 2,
			Banned: false,
			Address: model.Address{
				Street: "ulica", Number: "31", City: "Novi Sad", ZipCode: "21 000", Country: "Srbija",
			},
		},
	}

	for _, authority := range authorities {
		Driver.Create(&authority)
	}

	for _, user := range users {
		Driver.Create(&user)
	}

	fmt.Println("Created data.")

}
