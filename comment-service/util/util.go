package util

import (
	"comment-service/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"os"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetEnvDefault(envName string, defaultValue string) string {
	envValue := os.Getenv(envName);
	if envValue == "" {
		envValue = defaultValue;
	}
	return envValue;
}

func GetDocumentId(result *mongo.InsertOneResult) string {
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex()
	} else {
		return ""
	}
}

func GetUserId(request *http.Request) int {
	userServiceHost := GetEnvDefault("USER_SERVICE_HOST", "http://127.0.0.1:8081");
	client := &http.Client{}
	newRequest, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/users/userId", userServiceHost), nil)
	newRequest.Header.Add("Authorization", request.Header.Get("Authorization"))
	response, _ := client.Do(newRequest)

	var userIdDTO model.UserIdDTO
	_ = json.NewDecoder(response.Body).Decode(&userIdDTO)
	return userIdDTO.UserId
}

func UserCanComment(request *http.Request, drinkId uint64) bool {
	purchaseServiceHost := GetEnvDefault("PURCHASE_SERVICE_HOST", "http://127.0.0.1:8084");
	client := &http.Client{}
	userId := GetUserId(request);
	url := fmt.Sprintf("%s/api/purchases/user-comment-and-grade?user_id=", purchaseServiceHost) + strconv.FormatInt(int64(userId), 10) + "&drink_id=" + strconv.FormatUint(drinkId, 10);
	newRequest, _ := http.NewRequest("GET", url, nil)

	newRequest.Header.Add("Authorization", request.Header.Get("Authorization"))
	response, _ := client.Do(newRequest)

	var userGradeCheck model.UserCanGradeCheck
	_ = json.NewDecoder(response.Body).Decode(&userGradeCheck);
	return userGradeCheck.NumPurchases > 0;

}
