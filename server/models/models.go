package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Clients struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName    string             `json:"FirstName,omitempty"`
	LastName     string             `json:"LastName,omitempty"`
	PhoneNumber  string             `json:"PhoneNumber,omitempty"`
	EmailAddress string             `json:"EmailAddress,omitempty"`
	Description  string             `json:"Description,omitempty"`
	StateOfIssue string             `json:"StateOfIssue,omitempty"`
	FindHow      string             `json:"FindHow,omitempty"`
	SocialMedia  string             `json:"SocialMedia,omitempty"`
	Status       bool               `json:"status,omitempty"`
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
