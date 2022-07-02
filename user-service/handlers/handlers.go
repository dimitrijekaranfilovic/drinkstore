package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
	"user-service/database"
	"user-service/model"
	"user-service/repository"
)

var signingKey = []byte("signing_key")

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

func Authorize(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	queryParams := request.URL.Query()
	authority := queryParams.Get("authority")

	authorizationHeader := request.Header.Get("Authorization")
	tokenString := authorizationHeader[7:]
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})

	if err != nil || !token.Valid {
		fmt.Println("Ili ima gresku ili token ne valja")
		fmt.Println(err.Error())
		writer.WriteHeader(http.StatusForbidden)
		json.NewEncoder(writer).Encode(model.AuthorizationResponse{Allowed: false})
		return
	}

	if token.Claims.(*model.JwtClaims).Authority != authority {
		fmt.Println("Nije dobar autoritet")
		writer.WriteHeader(http.StatusForbidden)
		json.NewEncoder(writer).Encode(model.AuthorizationResponse{Allowed: false})
		return
	}
	fmt.Println("Sve ok")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(model.AuthorizationResponse{Allowed: true})
	return
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
			token, err2 := generateJwt(&user)
			if err2 != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(writer).Encode(model.Error{
					Message:       "Error creating token.",
					Timestamp:     time.Now(),
					Path:          "/api/users/authenticate",
					Status:        http.StatusInternalServerError,
					StatusMessage: "Internal Server Error.",
				})
			} else {
				//fmt.Println("Login uspjesan")
				json.NewEncoder(writer).Encode(model.AuthenticationResponseDTO{Jwt: token})
			}

		}
	}
}

func generateJwt(user *model.User) (string, error) {
	authority, err := repository.FindAuthorityById(user.AuthorityId)

	timestamp := time.Now()

	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Username
	claims["authority"] = authority.Name
	claims["iat"] = timestamp.Unix()
	claims["exp"] = timestamp.Add(time.Hour * 2).Unix()
	claims["userId"] = user.Id
	claims["iss"] = "ntp-user-service"

	signedToken, err2 := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}
	if err2 != nil {
		return "", err2
	}
	return signedToken, nil

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
