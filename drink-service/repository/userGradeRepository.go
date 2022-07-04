package repository

import (
	"drink-service/database"
	"drink-service/model"
	"errors"
	"strconv"
)

func DeleteUserGradeById(userGradeId uint) error {

	result := database.Driver.Delete(&model.UserGrade{}, userGradeId)
	if result.Error != nil {
		return errors.New("Error deleting user grade with id: " + strconv.FormatUint(uint64(userGradeId), 10) + ".")
	}

	return nil

}
