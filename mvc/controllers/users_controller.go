package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pratapsgit/gopractices/mvc/services"
	"github.com/pratapsgit/gopractices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		//Return the error as Bad Request
		apiErr := &utils.ApplicationError{
			Message:    "User Id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return

	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// Handle the error and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	//return user to the client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
