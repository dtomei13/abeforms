package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cases struct { // This is the claims by the clients.... SO IT WOULD BE BETTER IF IT WAS CALLED CLAIMS
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName      string             `json:"FirstName,omitempty" validate:"required"`
	LastName       string             `json:"LastName,omitempty" validate:"required"`
	EmailAddress   string             `json:"EmailAddress,omitempty" validate:"required"`
	Description    string             `json:"Description,omitempty" validate:"required"`
	StateOfIssue   string             `json:"StateOfIssue,omitempty" validate:"required"`
	LawyerAssigned string             `json:"LawyerAssigned,omitempty"`
	Status         bool               `json:"Status,omitempty"`
}
