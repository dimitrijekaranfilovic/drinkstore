package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
	"user-service/database"
	"user-service/model"
	"user-service/repository"
)

func RegisterUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	var userCreationDTO model.UserCreationDTO
	_ = json.NewDecoder(request.Body).Decode(&userCreationDTO)

	user := model.ToUser(&userCreationDTO)
	user.Password = database.HashPassword(userCreationDTO.Password)

	err := repository.CreateUser(&user)
	if err == nil {
		json.NewEncoder(writer).Encode(model.ToUserDTO(&user))
	} else {
		json.NewEncoder(writer).Encode(model.Error{
			Message:       err.Error(),
			Timestamp:     time.Now(),
			Path:          "/api/users/register",
			Status:        http.StatusBadRequest,
			StatusMessage: "Bad request.",
		})
	}

}

func Authenticate(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var authReqDTO model.AuthenticationRequestDTO
	_ = json.NewDecoder(request.Body).Decode(&authReqDTO)

	user, err := repository.FindUserByUsername(authReqDTO.Username)

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		createBadCredentialsError(writer)
	} else {
		hashComparison := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authReqDTO.Password))
		if hashComparison != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			createBadCredentialsError(writer)
		} else {
			//TODO: napravi token
			fmt.Println("Login uspjesan")
		}
	}
}

func createBadCredentialsError(writer http.ResponseWriter) {
	json.NewEncoder(writer).Encode(model.Error{
		Message:       "Bad credentials.",
		Timestamp:     time.Now(),
		Path:          "/api/users/authenticate",
		Status:        http.StatusUnauthorized,
		StatusMessage: "Unauthorized.",
	})
}

func BanUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	userId, _ := strconv.ParseUint(params["id"], 10, 64)

	user, err := repository.FindUserById(uint(userId))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(model.Error{
			Message:       err.Error(),
			Timestamp:     time.Now(),
			Path:          "/api/users/ban",
			Status:        http.StatusNotFound,
			StatusMessage: "Not found.",
		})
	} else {
		if user.AuthorityId == 1 {
			writer.WriteHeader(http.StatusForbidden)
			json.NewEncoder(writer).Encode(model.Error{
				Message:       "Admin accounts cannot be banned.",
				Timestamp:     time.Now(),
				Path:          "/api/users/ban",
				Status:        http.StatusForbidden,
				StatusMessage: "Forbidden.",
			})
		} else {
			user.Banned = true
			repository.UpdateUser(&user)
		}
	}
}
