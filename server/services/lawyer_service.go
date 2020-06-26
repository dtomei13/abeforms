package services

import (
	"github.com/austinlhx/server/domain"
	"github.com/austinlhx/server/models"
	"github.com/austinlhx/server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetLawyer(email string, password string)(*models.Lawyers, *utils.ApplicationError){
	return domain.GetLawyer(email, password)
}

func AddLawyer(lawyerInfo models.Lawyers) *utils.ApplicationError {
	return domain.AddLawyer(lawyerInfo)
}

func GetCase()([]primitive.M, *utils.ApplicationError){
	return domain.GetCase()
}

func GetEachCase(user models.Lawyers)([]primitive.M, *utils.ApplicationError){
	return domain.GetEachCase(user)
}

func TakeCase(caseID string, user models.Lawyers)*utils.ApplicationError{
	return domain.TakeCase(caseID, user)
}

func CreateMeeting(zoomInfo utils.ZoomMeeting, user models.Lawyers)*utils.ApplicationError{
	return domain.CreateMeeting(zoomInfo, user)
}