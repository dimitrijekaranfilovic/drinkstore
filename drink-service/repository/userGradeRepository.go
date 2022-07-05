package repository

import (
	"drink-service/database"
	"drink-service/model"
	"errors"
	"strconv"
)

func CreateGrade(grade *model.UserGrade) {
	database.Driver.Create(grade)
}

func DeleteUserGrade(drinkId uint, userGradeId uint, userId uint) error {

	result := database.Driver.Where("drink_id = ? AND id = ? AND user_id = ?", drinkId, userGradeId, userId).Delete(&model.UserGrade{})
	if result.Error != nil {
		return errors.New("Error deleting user grade with id: " + strconv.FormatUint(uint64(userGradeId), 10) + ".")
	}
	if result.RowsAffected == 0 {
		return errors.New("Desired grade does not exist.")
	}
	return nil
}

func FindUserGrade(drinkId uint, userGradeId uint, userId uint) (model.UserGrade, error) {
	var userGrade model.UserGrade
	database.Driver.First(&userGrade, "drink_id = ? AND id = ? AND user_id = ?", drinkId, userGradeId, userId)
	if userGrade.Id == 0 {
		return userGrade, errors.New("Desired user grade does not exist.")
	}
	return userGrade, nil
}

func CheckForGradeForUserAndDrink(userId uint, drinkId uint) error {
	var userGrade model.UserGrade
	result := database.Driver.Where("drink_id = ? AND user_id = ?", drinkId, userId).First(&userGrade)
	if result.Error != nil {
		return nil
	} else {
		return errors.New("User grade for user with id: " + strconv.FormatUint(uint64(userId), 10) + " and drink with id: " + strconv.FormatUint(uint64(drinkId), 10) + " already exists.")
	}
}

func UpdateUserGrade(userGrade *model.UserGrade) {
	database.Driver.Save(userGrade)
}

func GetGradeForDrinkAndUser(userId uint, drinkId uint) (model.UserGrade, error) {
	var userGrade model.UserGrade
	database.Driver.First(&userGrade, "user_id = ? AND drink_id = ?", userId, drinkId)
	if userGrade.Id == 0 {
		return userGrade, errors.New("Grade does not exist.")
	}
	return userGrade, nil
}
