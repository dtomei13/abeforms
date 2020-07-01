package domain

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"../database"
	"../models"
	"../utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUnassignedCase(user models.Clients) ([]primitive.M, *utils.ApplicationError) {
	collection := database.ClaimsCollection
	filter := &bson.M{"emailaddress": user.EmailAddress}
	data, err := collection.Find(context.TODO(), filter)
	var cases []primitive.M

	for data.Next(context.Background()) {
		var lawyerEmail bson.M
		e := data.Decode(&lawyerEmail)
		if e != nil {
			log.Fatal(e)
		}
		cases = append(cases, lawyerEmail)
	}

	if err != nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Getting cases from DB failed"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}

	data.Close(context.Background())

	var unassignedCases []primitive.M
	for _, b := range cases {
		if b[("lawyerassigned")] == "" {
			unassignedCases = append(unassignedCases, b)
		}
	}
	return unassignedCases, nil
}

func GetAssignedCase(user models.Clients) ([]primitive.M, *utils.ApplicationError) {
	collection := database.ClaimsCollection
	filter := &bson.M{"emailaddress": user.EmailAddress}
	data, err := collection.Find(context.TODO(), filter)
	var cases []primitive.M
	for data.Next(context.Background()) {
		var lawyerEmail bson.M
		e := data.Decode(&lawyerEmail)
		if e != nil {
			log.Fatal(e)
		}
		cases = append(cases, lawyerEmail)
	}
	if err != nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Getting cases from DB failed"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}

	data.Close(context.Background())

	var assignedCases []primitive.M
	for _, b := range cases {
		if b[("lawyerassigned")] != "" {
			assignedCases = append(assignedCases, b)
		}
	}
	return assignedCases, nil
}
