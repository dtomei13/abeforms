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
)

func GetLawyer(email string, password string) (*models.Lawyers, *utils.ApplicationError) {
	fmt.Println(email)
	collection := database.LawyerCollection //this works fine
	filter := &bson.M{
		"emailaddress": email,
	}
	var lawyer *models.Lawyers //Below is not working
	err := collection.FindOne(context.TODO(), filter).Decode(&lawyer)
	if err != nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Either the emailaddress or password was incorrect"),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	//check if password is correct
	if password != lawyer.Password {
		log.Println("Password Incorrect")
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Password was incorrect!"),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	//lawyer found
	log.Println("Lawyer Found!")

	return lawyer, nil
}

func AddLawyer(lawyerInfo models.Lawyers) *utils.ApplicationError {
	//TODO: Check if email is a duplicate
	collection := database.LawyerCollection
	if _, err := collection.InsertOne(context.Background(), lawyerInfo); err != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Adding to DB failed"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}
	log.Println("Lawyer added to DB!")

	return nil //no error!
}
