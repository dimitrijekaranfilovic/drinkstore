package model

import "time"

type UserDTO struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

func ToUserDTO(user *User) UserDTO {
	return UserDTO{Id: user.Id, Username: user.Username}
}

type AuthenticationRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticationResponseDTO struct {
	Jwt string `json:"jwt"`
}

type Error struct {
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
	Path          string    `json:"path"`
	Status        int       `json:"status"`
	StatusMessage string    `json:"statusMessage"`
}

type UserCreationDTO struct {
	Username  string  `json:"username" gorm:"unique"`
	Password  string  `json:"password"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Address   Address `json:"address" gorm:"embedded"`
}
