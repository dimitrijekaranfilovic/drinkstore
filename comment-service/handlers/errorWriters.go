package handlers

import (
	"comment-service/model"
	"encoding/json"
	"net/http"
	"time"
)

func writeInternalServerError(writer http.ResponseWriter, request *http.Request, err error) {
	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(model.Error{
		Message:       err.Error(),
		Timestamp:     time.Now(),
		Path:          request.RequestURI,
		Status:        http.StatusInternalServerError,
		StatusMessage: "Internal server error.",
	})
}

func writeBadRequest(writer http.ResponseWriter, request *http.Request, err error) {
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(model.Error{
		Message:       err.Error(),
		Timestamp:     time.Now(),
		Path:          request.RequestURI,
		Status:        http.StatusBadRequest,
		StatusMessage: "Bad request.",
	})
}

func writeBadCredentials(writer http.ResponseWriter) {
	json.NewEncoder(writer).Encode(model.Error{
		Message:       "Bad credentials.",
		Timestamp:     time.Now(),
		Path:          "/api/users/authenticate",
		Status:        http.StatusUnauthorized,
		StatusMessage: "Unauthorized.",
	})
}

func writeNotFound(writer http.ResponseWriter, request *http.Request, err error) {
	writer.WriteHeader(http.StatusNotFound)
	json.NewEncoder(writer).Encode(model.Error{
		Message:       err.Error(),
		Timestamp:     time.Now(),
		Path:          request.RequestURI,
		Status:        http.StatusNotFound,
		StatusMessage: "Not found.",
	})
}
