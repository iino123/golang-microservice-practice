package controllers

import (
	"encoding/json"
	"github.com/iino123/golang-microservice-practice/mvc/services"
	"github.com/iino123/golang-microservice-practice/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr2 := services.GetUser(userId)
	if apiErr2 != nil {
		jsonValue, _ := json.Marshal(apiErr2)
		resp.WriteHeader(apiErr2.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
