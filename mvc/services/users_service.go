package services

import (
	"github.com/pratapsgit/gopractices/mvc/domain"
	"github.com/pratapsgit/gopractices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
