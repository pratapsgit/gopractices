package domain

import (
	"fmt"
	"net/http"

	"github.com/pratapsgit/gopractices/mvc/utils"
)

var (
	users = map[int64]*User{
		124: {124, "George", "Derek", "george.derek@myorg.edu"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
