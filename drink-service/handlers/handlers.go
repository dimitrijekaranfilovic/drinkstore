package handlers

import (
	"drink-service/model"
	"drink-service/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func CreateDrink(writer http.ResponseWriter, request *http.Request) {
	//kreiram pice, vratim id
	writer.Header().Set("Content-Type", "application/json")

	var drinkCreationDTO model.DrinkCreationDTO
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

}

func UpdateUserGrade(writer http.ResponseWriter, request *http.Request) {

}

func DeleteUserGrade(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	userGradeId, _ := strconv.ParseUint(params["id"], 10, 64)

	err := repository.DeleteUserGradeById(uint(userGradeId))

	if err != nil {
		writeInternalServerError(writer, request, err)
	} else {
		writer.WriteHeader(http.StatusNoContent)

	}

}
