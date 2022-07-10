package util

import (
	"drink-service/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUserId(request *http.Request) int {
	client := &http.Client{}
	newRequest, _ := http.NewRequest("GET", "http://127.0.0.1:8081/api/users/userId", nil)
	newRequest.Header.Add("Authorization", request.Header.Get("Authorization"))
	response, _ := client.Do(newRequest)

	var userIdDTO model.UserIdDTO
	_ = json.NewDecoder(response.Body).Decode(&userIdDTO)
	return userIdDTO.UserId
}


func UserCanGrade(request *http.Request, drinkId uint64) bool {
	client := &http.Client{}

	userId := GetUserId(request);
	url := "http://127.0.0.1:8084/api/purchases/user-comment-and-grade?user_id=" + strconv.FormatInt(int64(userId), 10) + "&drink_id=" + strconv.FormatUint(drinkId, 10);
	newRequest, _ := http.NewRequest("GET", url, nil)

	newRequest.Header.Add("Authorization", request.Header.Get("Authorization"))
	response, _ := client.Do(newRequest)

	var userGradeCheck model.UserCanGradeCheck
	_ = json.NewDecoder(response.Body).Decode(&userGradeCheck);
	return userGradeCheck.NumPurchases > 0;

}