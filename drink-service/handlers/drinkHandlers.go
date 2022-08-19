package handlers

import (
	"drink-service/model"
	"drink-service/repository"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func GetSingleDrink(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	drinkId, _ := strconv.ParseUint(params["id"], 10, 64)

	drink, err := repository.FindDrinkById(uint(drinkId))
	if err != nil {
		writeNotFound(writer, request, err)
		return
	}
	json.NewEncoder(writer).Encode(drink)
}
func GetUnfilteredDrinks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(repository.GetAllDrinks())

}

func GetDrinks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	queryParams := request.URL.Query()

	categories := strings.Split(queryParams.Get("categories"), ",")
	volumeLabels := strings.Split(queryParams.Get("volumeLabels"), ",")
	query := queryParams.Get("query")
	gradeFrom := queryParams.Get("gradeFrom")
	gradeTo := queryParams.Get("gradeTo")

	page := queryParams.Get("page")
	size := queryParams.Get("size")
	sortCriteria := queryParams.Get("sortCriteria")
	sortDirection := queryParams.Get("sortDirection")

	gradeFromParsed, _ := strconv.ParseFloat(gradeFrom, 64)
	gradeToParsed, _ := strconv.ParseFloat(gradeTo, 64)
	pageParsed, _ := strconv.ParseUint(page, 10, 64)
	sizeParsed, _ := strconv.ParseUint(size, 10, 64)

	drinks, totalItems := repository.GetDrinks(categories, volumeLabels, query, gradeFromParsed, gradeToParsed, pageParsed, sizeParsed, sortCriteria, sortDirection)

	var totalPages int64

	totalPages = totalItems / int64(sizeParsed)

	if totalItems%int64(sizeParsed) != 0 {
		totalPages += 1
	}

	drinkPage := model.DrinkPage{
		Drinks:        drinks,
		Page:          pageParsed,
		SortCriteria:  sortCriteria,
		SortDirection: sortDirection,
		TotalPages:    totalPages,
	}

	json.NewEncoder(writer).Encode(drinkPage)
}

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

func updateDrinkAverageGrade(drinkId uint) {
	repository.UpdateDrinkAverageGrade(drinkId)
}

func GetSingleImage(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	imageName, _ := params["imageName"]
	imageType := strings.ToLower(strings.Split(imageName, ".")[1])
	writer.Header().Set("Content-Type", "image/"+imageType)
	http.ServeFile(writer, request, "images/"+imageName)
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
			imagePath := "images/" + strings.ReplaceAll(handler.Filename, " ", "-")
			file, err3 := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)

			if err3 != nil {
				writeInternalServerError(writer, request, err3)

			} else {
				_, _ = io.Copy(file, imageFile)
				if drink.ImagePath != "" {
					err4 := os.Remove(drink.ImagePath)
					if err4 != nil {
						writeInternalServerError(writer, request, err4)
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
		drink.Price = drinkUpdateDTO.Price
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
