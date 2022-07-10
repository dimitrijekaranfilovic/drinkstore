package repository

import (
	"drink-service/database"
	"drink-service/model"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func GetAllDrinks() []model.Drink {
	var drinks []model.Drink
	database.Driver.Find(&drinks)
	return drinks
}

func GetDrinks(categories []string, volumeLabels []string, query string, gradeFrom float64, gradeTo float64, page uint64, size uint64, sortCriteria string, sortDirection string) ([]model.Drink, int64) {
	var drinks, totalDrinks []model.Drink
	var totalItems int64

	query = "%" + strings.ToLower(query) + "%"
	database.Driver.
		Where(
			"category IN ? AND "+
				"volume_label IN ? AND "+
				"(lower(name) LIKE ? OR lower(description) LIKE ?) AND "+
				"average_grade >= ? "+
				"AND average_grade <= ? ",
			categories,
			volumeLabels,
			query,
			query,
			gradeFrom,
			gradeTo).
		Offset(int(page * size)).
		Limit(int(size)).
		Order(sortCriteria + " " + sortDirection).
		Find(&drinks)

	database.Driver.
		Where(
			"category IN ? AND "+
				"volume_label IN ? AND "+
				"(lower(name) LIKE ? OR lower(description) LIKE ?) AND "+
				"average_grade >= ? "+
				"AND average_grade <= ? ",
			categories,
			volumeLabels,
			query,
			query,
			gradeFrom,
			gradeTo).
		Find(&totalDrinks).Count(&totalItems)

	fmt.Println("Total items: " + strconv.FormatInt(totalItems, 10))
	return drinks, totalItems
}

func DeleteDrinkById(drinkId uint) error {
	result1 := database.Driver.Where("drink_id = ?", drinkId).Delete(&model.UserGrade{})
	if result1.Error != nil {
		return errors.New("Error deleting user grades for drink with id: " + strconv.FormatUint(uint64(drinkId), 10) + ".")
	}

	result2 := database.Driver.Delete(&model.Drink{}, drinkId)
	if result2.Error != nil {
		return errors.New("Error deleting drink with id: " + strconv.FormatUint(uint64(drinkId), 10) + ".")
	}

	if result2.RowsAffected == 0 {
		return errors.New("Drink with id: " + strconv.FormatUint(uint64(drinkId), 10) + " does not exist.")

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

func UpdateDrinkAverageGrade(drinkId uint) {
	subQuery := database.Driver.Select("avg(grade)").Where("drink_id = ?", drinkId).Table("user_grades")
	database.Driver.Model(&model.Drink{}).Where("id = ?", drinkId).Update("average_grade", subQuery)
}
