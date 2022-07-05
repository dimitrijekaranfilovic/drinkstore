package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"user-service/database"
	"user-service/model"
	"user-service/repository"
)

var signingKey = []byte("signing_key")

func RegisterUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userCreationDTO model.UserCreationDTO
	_ = json.NewDecoder(request.Body).Decode(&userCreationDTO)

	user := model.ToUser(&userCreationDTO)
	user.Password = database.HashPassword(userCreationDTO.Password)

	err := repository.CreateUser(&user)
	if err == nil {
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(model.ToUserDTO(&user))
	} else {
		writeBadRequest(writer, request, err)
	}

}

func extractJWT(request *http.Request) (*jwt.Token, error) {
	authorizationHeader := request.Header.Get("Authorization")
	if len(authorizationHeader) == 0 {
		return nil, errors.New("No Authorization header found.")
	}
	if !strings.Contains(authorizationHeader, "Bearer") {
		return nil, errors.New("Authorization header is not 'Bearer'.")
	}
	tokenString := authorizationHeader[7:]
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})

	return token, err
}

func Authorize(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	queryParams := request.URL.Query()
	authority := queryParams.Get("authority")

	token, err := extractJWT(request)

	if err != nil || !token.Valid {
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
		writeBadCredentials(writer)
	} else {
		hashComparison := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authReqDTO.Password))
		if hashComparison != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			writeBadCredentials(writer)
		} else {
			//TODO: napravi token
			token, err2 := generateJwt(&user)
			if err2 != nil {
				writeInternalServerError(writer, request, err2)
			} else {
				writer.WriteHeader(http.StatusOK)
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
	claims["exp"] = timestamp.Add(time.Hour * 100).Unix()
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

func BanUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	userId, _ := strconv.ParseUint(params["id"], 10, 64)

	user, err := repository.FindUserById(uint(userId))
	if err != nil {
		writeNotFound(writer, request, err)
	} else {
		if user.AuthorityId == 1 {
			writer.WriteHeader(http.StatusForbidden)
			json.NewEncoder(writer).Encode(model.Error{
				Message:       "Admin accounts cannot be banned.",
				Timestamp:     time.Now(),
				Path:          request.RequestURI,
				Status:        http.StatusForbidden,
				StatusMessage: "Forbidden.",
			})
		} else {
			user.Banned = true
			repository.UpdateUser(&user)
		}
	}
}

func GetUserIdFromJWT(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	token, err := extractJWT(request)
	if err != nil || !token.Valid {
		fmt.Println(err.Error())
		writer.WriteHeader(http.StatusForbidden)
		json.NewEncoder(writer).Encode(model.UserIdDTO{UserId: -1})
	} else {
		writer.WriteHeader(http.StatusOK)
		userId := token.Claims.(*model.JwtClaims).UserId
		json.NewEncoder(writer).Encode(model.UserIdDTO{UserId: int(userId)})
	}

}
