package repository

import (
	"errors"
	"strconv"
	"user-service/database"
	"user-service/model"
)

func GetAllUsers() ([]model.User, error) {

	var users []model.User
	result := database.Driver.Find(&users)
	return users, result.Error
}

func CreateUser(user *model.User) error {
	database.Driver.Create(user)
	if user.Id == 0 {
		return errors.New("User with username: " + user.Username + " already exists.")
	}
	return nil
}

func FindUserByUsername(username string) (model.User, error) {
	var user model.User
	result := database.Driver.First(&user, "username = ? AND banned = false", username)
	return user, result.Error
}

func FindUserById(id uint) (model.User, error) {
	var user model.User
	database.Driver.First(&user, "id = ?", id)
	if user.Id == 0 {
		return user, errors.New("User with id: " + strconv.FormatUint(uint64(id), 10) + " does not exist.")

	}
	return user, nil
}

func UpdateUser(user *model.User) {
	database.Driver.Save(user)
}
