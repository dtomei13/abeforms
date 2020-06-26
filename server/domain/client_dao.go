package domain

import (
	"context"
	"fmt"
	"github.com/austinlhx/server/database"
	"github.com/austinlhx/server/models"
	"github.com/austinlhx/server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
)

func AddClient(clientInfo models.Clients) *utils.ApplicationError {
	//TODO: Check if email is a duplicate
	collection := database.ClientCollection
	if _, err := collection.InsertOne(context.Background(), clientInfo); err != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Adding to DB failed"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}
	log.Println("Client added to DB!")


	return nil
}

func SendEmail() *utils.ApplicationError{

	collection := database.LawyerCollection

	data, err := collection.Find(context.Background(), bson.D{{}})
	//TODO: try and figure out how to filter only emails
	if err != nil{
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Finding the data from DB Failed!"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}
	var lawyerEmails []primitive.M
	for data.Next(context.Background()) {
		var lawyerEmail bson.M
		err := data.Decode(&lawyerEmail)
		if err != nil {
			return &utils.ApplicationError{
				Message:    fmt.Sprintf("Decoding lawyer data failed"),
				StatusCode: http.StatusInternalServerError,
				Code:       "internal_error",
			}
		}
		lawyerEmails = append(lawyerEmails, lawyerEmail)
	}

	if err := data.Err(); err != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Reading Client Data failed"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}

	data.Close(context.Background())
	var allEmails [] string
	for _, b := range lawyerEmails{
		allEmails = append(allEmails, b["emailaddress"].(string))
	}
	log.Println(allEmails)

	auth := smtp.PlainAuth("", "austin.abe.legal@gmail.com", "Password", "smtp.gmail.com")
	to := allEmails
	msg := []byte(
		"To: " + "austin.abe.legal@gmail.com" + "\r\n" +
		"Subject: Abe Consult Alert\r\n" +
		"\r\n" +
		"New Client Alert! View now on our website :)")
	error := smtp.SendMail("smtp.gmail.com:587", auth, "austin.abe.legal@gmail.com", to, msg)
	if error != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Sending Emails failed"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	} else {
		fmt.Println("Sent Email to " + strconv.Itoa(len(lawyerEmails)) + " Lawyers!")
	}
	return nil
}