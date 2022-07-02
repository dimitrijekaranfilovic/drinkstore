package model

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type ApiError interface {
	ApiError() (string, time.Time, string, int, string)
}

type Authority struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"authority"`
}

type User struct {
	Id          uint    `json:"id" gorm:"primaryKey"`
	Username    string  `json:"username" gorm:"unique"`
	Password    string  `json:"password"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	AuthorityId uint    `json:"authority"`
	Address     Address `json:"address" gorm:"embedded"`
	Banned      bool    `json:"banned"`
}

type Address struct {
	Street  string `json:"street"`
	Number  string `json:"number"`
	City    string `json:"city"`
	ZipCode string `json:"zipCode"`
	Country string `json:"country"`
}

func ToUser(dto *UserCreationDTO) User {
	return User{
		Username:    dto.Username,
		Password:    dto.Password,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Address:     dto.Address,
		AuthorityId: 2, //USER authority,
		Banned:      false,
	}
}

type JwtClaims struct {
	Authority string `json:"authority"`
	jwt.StandardClaims
}
