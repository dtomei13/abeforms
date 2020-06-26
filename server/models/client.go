package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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