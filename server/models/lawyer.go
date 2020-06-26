package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lawyers struct {
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
