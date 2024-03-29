package util

import (
	"comment-service/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDocumentId(result *mongo.InsertOneResult) string {
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex()
	} else {
		return ""
	}
}

func GetUserId(request *http.Request) int {
	userServiceHost := os.Getenv("USER_SERVICE_HOST")
	if userServiceHost == "" {
		userServiceHost = "127.0.0.1"
	}
	client := &http.Client{}
	newRequest, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:8081/api/users/userId", userServiceHost), nil)
	newRequest.Header.Add("Authorization", request.Header.Get("Authorization"))
	response, _ := client.Do(newRequest)

	var userIdDTO model.UserIdDTO
	_ = json.NewDecoder(response.Body).Decode(&userIdDTO)
	return userIdDTO.UserId
}

func UserCanComment(request *http.Request, drinkId uint64) bool {
	purchaseServiceHost := os.Getenv("PURCHASE_SERVICE_HOST")
	if purchaseServiceHost == "" {
		purchaseServiceHost = "127.0.0.1"
	}

	client := &http.Client{}
	userId := GetUserId(request);
	url := fmt.Sprintf("http://%s:8084/api/purchases/user-comment-and-grade?user_id=", purchaseServiceHost) + strconv.FormatInt(int64(userId), 10) + "&drink_id=" + strconv.FormatUint(drinkId, 10);
	//fmt.Println("url: " + url)
	newRequest, _ := http.NewRequest("GET", url, nil)

	newRequest.Header.Add("Authorization", request.Header.Get("Authorization"))
	response, _ := client.Do(newRequest)

	var userGradeCheck model.UserCanGradeCheck
	_ = json.NewDecoder(response.Body).Decode(&userGradeCheck);
	return userGradeCheck.NumPurchases > 0;

}
