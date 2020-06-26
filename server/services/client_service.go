package services

import (
	"github.com/austinlhx/server/domain"
	"github.com/austinlhx/server/models"
	"github.com/austinlhx/server/utils"
)

func AddClient(clientInfo models.Clients) *utils.ApplicationError {
	return domain.AddClient(clientInfo)
}

func SendEmail() *utils.ApplicationError{
	return domain.SendEmail()
}