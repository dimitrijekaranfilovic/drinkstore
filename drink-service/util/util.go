package util

import (
	"drink-service/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"drink-service/database"
)

func GetUserId(request *http.Request) int {
	userServiceHost := database.GetEnvDefault("USER_SERVICE_HOST", "http://127.0.0.1:8081");
	client := &http.Client{}
	newRequest, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/users/userId", userServiceHost), nil)
	newRequest.Header.Add("Authorization", request.Header.Get("Authorization"))
	response, _ := client.Do(newRequest)

	var userIdDTO model.UserIdDTO
	_ = json.NewDecoder(response.Body).Decode(&userIdDTO)
	return userIdDTO.UserId
}


func UserCanGrade(request *http.Request, drinkId uint64) bool {
	
	purchaseServiceHost := database.GetEnvDefault("PURCHASE_SERVICE_HOST", "http://127.0.0.1:8084");

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