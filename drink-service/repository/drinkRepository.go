package repository

import (
	"drink-service/database"
	"drink-service/model"
	"errors"
	"strconv"
)

func DeleteDrinkById(drinkId uint) error {
	result1 := database.Driver.Where("drink_id = ?", drinkId).Delete(&model.UserGrade{})
	if result1.Error != nil {
		return errors.New("Error deleting user grades for drink with id: " + strconv.FormatUint(uint64(drinkId), 10) + ".")
	}

	result2 := database.Driver.Delete(&model.Drink{}, drinkId)
	if result2.Error != nil {
		return errors.New("Error deleting drink with id: " + strconv.FormatUint(uint64(drinkId), 10) + ".")
	}

	return nil
}

func CreateDrink(drink *model.Drink) error {
	database.Driver.Create(drink)
	if drink.Id == 0 {
		return errors.New("Drink with name: " + drink.Name + " already exists.")
	}
	return nil
}

func FindDrinkById(drinkId uint) (model.Drink, error) {
	var drink model.Drink
	database.Driver.First(&drink, "id = ?", drinkId)
	if drink.Id == 0 {
		return drink, errors.New("Drink with id: " + strconv.FormatUint(uint64(drinkId), 10) + " does not exist.")
	}
	return drink, nil
}

func UpdateDrink(drink *model.Drink) {
	database.Driver.Save(drink)
}
