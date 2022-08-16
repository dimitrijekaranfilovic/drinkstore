package handlers

import (
	"drink-service/model"
	"drink-service/repository"
	"drink-service/util"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUserGrade(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	userId := util.GetUserId(request)
	if userId == -1 {
		writeBadRequest(writer, request, errors.New("You cannot grade this drink."))
		return
	}
	fmt.Println("User id: " + strconv.FormatInt(int64(userId), 10))

	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["id"], 10, 64)

	_, err := repository.FindDrinkById(uint(drinkId))
	if err != nil {
		writeBadRequest(writer, request, err)
		return
	}

	hasOrdered := util.UserCanGrade(request, drinkId);
	if hasOrdered {
		var userGradeDTO model.UserGradeDTO
		_ = json.NewDecoder(request.Body).Decode(&userGradeDTO)
		err := repository.CheckForGradeForUserAndDrink(uint(userId), uint(drinkId))
		if err != nil {
			writeBadRequest(writer, request, errors.New("You have already graded this drink."))
			return
		}

		userGrade := model.UserGrade{
			Grade:   userGradeDTO.Grade,
			DrinkId: uint(drinkId),
			UserId:  uint(userId)}

		repository.CreateGrade(&userGrade)
		updateDrinkAverageGrade(uint(drinkId))
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(userGrade)

	} else {
		message := "You cannot grade drink with id: " + strconv.FormatInt(int64(drinkId), 10) + " because you haven't purchased any."
		writeBadRequest(writer, request, errors.New(message))
	}
}

func UpdateUserGrade(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["drinkId"], 10, 64)
	userGradeId, _ := strconv.ParseUint(params["gradeId"], 10, 64)
	userId := util.GetUserId(request)

	userGrade, err := repository.FindUserGrade(uint(drinkId), uint(userGradeId), uint(userId))
	if err != nil {
		writeNotFound(writer, request, err)
	} else {
		var userGradeDTO model.UserGradeDTO
		_ = json.NewDecoder(request.Body).Decode(&userGradeDTO)

		userGrade.Grade = userGradeDTO.Grade
		repository.UpdateUserGrade(&userGrade)
		updateDrinkAverageGrade(uint(drinkId))

		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(userGrade)

	}

}

func DeleteUserGrade(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["drinkId"], 10, 64)
	userGradeId, _ := strconv.ParseUint(params["gradeId"], 10, 64)
	userId := util.GetUserId(request)

	err := repository.DeleteUserGrade(uint(drinkId), uint(userGradeId), uint(userId))

	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		updateDrinkAverageGrade(uint(drinkId))
		writer.WriteHeader(http.StatusNoContent)

	}

}

func CheckDrinkGradeForUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["id"], 10, 64)
	userId := util.GetUserId(request)

	userGrade, err := repository.GetGradeForDrinkAndUser(uint(userId), uint(drinkId))
	if err != nil {
		json.NewEncoder(writer).Encode(model.GradeCheckDTO{GradeExists: false, GradeValue: -1, GradeId: userGrade.Id})
	} else {
		json.NewEncoder(writer).Encode(model.GradeCheckDTO{GradeExists: true, GradeValue: int16(userGrade.Grade), GradeId: userGrade.Id})
	}
}
