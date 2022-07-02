package repository

import (
	"errors"
	"strconv"
	"user-service/database"
	"user-service/model"
)

func FindAuthorityById(id uint) (model.Authority, error) {
	var authority model.Authority
	database.Driver.First(&authority, "id = ?", id)
	if authority.Id == 0 {
		return authority, errors.New("Authority with id: " + strconv.FormatUint(uint64(id), 10) + " does not exist.")

	}
	return authority, nil
}
