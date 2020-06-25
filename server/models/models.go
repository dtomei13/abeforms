package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Clients struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName      string             `json:"FirstName,omitempty" validate:"required"`
	LastName       string             `json:"LastName,omitempty" validate:"required"`
	PhoneNumber    string             `json:"PhoneNumber,omitempty" validate:"required"`
	EmailAddress   string             `json:"EmailAddress,omitempty" validate:"required"`
	Description    string             `json:"Description,omitempty" validate:"required"`
	StateOfIssue   string             `json:"StateOfIssue,omitempty" validate:"required"`
	FindHow        string             `json:"FindHow,omitempty"`
	SocialMedia    string             `json:"SocialMedia,omitempty"`
	LawyerAssigned string             `json:"LawyerAssigned,omitempty"`
	Status         bool               `json:"Status,omitempty"`
}

type LawyerSignUp struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName      string             `json:"FirstName,omitempty"`
	LastName       string             `json:"LastName,omitempty"`
	PhoneNumber    string             `json:"PhoneNumber,omitempty"`
	EmailAddress   string             `json:"EmailAddress,omitempty"`
	StateOfLicense string             `json:"StateOfLicense,omitempty"`
	Expertise      string             `json:"Expertise,omitempty"`
	Password       string             `json:"Password,omitempty"`
	RetypePassword string             `json:"RetypePassword,omitempty"`
}

type LawyerSignIn struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmailAddress string             `json:"EmailAddress,omitempty"`
	Password     string             `json:"Password,omitempty"`
}

type TokenDetails struct {
	EmailAddress string `json:"EmailAddress,omitempty"`
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type ZoomMeeting struct {
	ClientEmail  string `json:"clientEmail,omitempty"`
	UserEmail    string `json:"userEmail,omitempty"`
	SelectedTime string `json:"selectedTime,omitempty"`
}
