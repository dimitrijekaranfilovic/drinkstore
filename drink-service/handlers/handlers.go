package handlers

import (
	"drink-service/model"
	"drink-service/repository"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func CreateDrink(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var drinkCreationDTO model.DrinkCreateUpdateDTO
	_ = json.NewDecoder(request.Body).Decode(&drinkCreationDTO)

	drink := model.ToDrinkFromDrinkCreationDTO(&drinkCreationDTO)

	err := repository.CreateDrink(&drink)
	if err == nil {
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(drink)
	} else {
		writeBadRequest(writer, request, err)
	}
}

func CreateUpdateDrinkImage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["id"], 10, 64)

	drink, err := repository.FindDrinkById(uint(drinkId))

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(model.Error{
			Message:       err.Error(),
			Timestamp:     time.Now(),
			Path:          request.RequestURI,
			Status:        http.StatusNotFound,
			StatusMessage: "Not found.",
		})
	} else {
		imageFile, handler, err2 := request.FormFile("image")
		if err2 != nil {
			writeInternalServerError(writer, request, err2)
		} else {
			imagePath := "images/" + handler.Filename
			file, err3 := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)

			if err3 != nil {
				writeInternalServerError(writer, request, err3)

			} else {
				_, _ = io.Copy(file, imageFile)
				if drink.ImagePath != "" {
					err4 := os.Remove(drink.ImagePath)
					if err4 != nil {

					}
				}
				drink.ImagePath = imagePath
				repository.UpdateDrink(&drink)
			}
			file.Close()
		}
		imageFile.Close()
	}
}

func UpdateDrink(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["id"], 10, 64)

	drink, err := repository.FindDrinkById(uint(drinkId))
	if err != nil {
		writeNotFound(writer, request, err)
	} else {
		var drinkUpdateDTO model.DrinkCreateUpdateDTO
		_ = json.NewDecoder(request.Body).Decode(&drinkUpdateDTO)
		drink.Name = drinkUpdateDTO.Name
		drink.Description = drinkUpdateDTO.Description
		drink.Volume = drinkUpdateDTO.Volume
		drink.VolumeLabel = drinkUpdateDTO.VolumeLabel
		drink.Category = drinkUpdateDTO.Category
		repository.UpdateDrink(&drink)
		writer.WriteHeader(http.StatusNoContent)
	}
}

func DeleteDrink(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["id"], 10, 64)

	err := repository.DeleteDrinkById(uint(drinkId))
	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		writer.WriteHeader(http.StatusNoContent)

	}
}

func CreateUserGrade(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	userId := getUserId(request)
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

	//TODO: ovdje pozovi purchase service
	hasOrdered := true
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
		writer.WriteHeader(http.StatusCreated)

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
	userId := getUserId(request)

	userGrade, err := repository.FindUserGrade(uint(drinkId), uint(userGradeId), uint(userId))
	if err != nil {
		writeNotFound(writer, request, err)
		return
	}

	var userGradeDTO model.UserGradeDTO
	_ = json.NewDecoder(request.Body).Decode(&userGradeDTO)

	userGrade.Grade = userGradeDTO.Grade
	repository.UpdateUserGrade(&userGrade)

	writer.WriteHeader(http.StatusNoContent)

}

func DeleteUserGrade(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["drinkId"], 10, 64)
	userGradeId, _ := strconv.ParseUint(params["gradeId"], 10, 64)
	userId := getUserId(request)

	err := repository.DeleteUserGrade(uint(drinkId), uint(userGradeId), uint(userId))

	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		writer.WriteHeader(http.StatusNoContent)

	}

}

func getUserId(request *http.Request) int {
	client := &http.Client{}
	newRequest, _ := http.NewRequest("GET", "http://127.0.0.1:8081/api/users/userId", nil)
	newRequest.Header.Add("Authorization", request.Header.Get("Authorization"))
	response, _ := client.Do(newRequest)

	var userIdDTO model.UserIdDTO
	_ = json.NewDecoder(response.Body).Decode(&userIdDTO)
	return userIdDTO.UserId
}
