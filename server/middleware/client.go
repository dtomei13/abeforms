package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strconv"

	"../models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateClientsInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var client models.Clients
	_ = json.NewDecoder(r.Body).Decode(&client)
	insertOneClient(client)

	json.NewEncoder(w).Encode(client)
	lawyersPrimitive := getInfo(lawyerCollection)
	var lawyersEmails []string

	for _, b := range lawyersPrimitive {
		lawyersEmails = append(lawyersEmails, b[("emailaddress")].(string))
	}

	clientInfo := map[string]string{
		"FirstName":    client.FirstName,
		"LastName":     client.LastName,
		"EmailAddress": client.EmailAddress,
		"PhoneNumber":  client.PhoneNumber,
		"Location":     client.StateOfIssue,
		"Description":  client.Description,
		"FindHow":      client.FindHow,
		"SocialMedia":  client.SocialMedia,
	}

	fmt.Println(lawyersEmails)
	fmt.Println(clientInfo)

	sendEmails(lawyersEmails, clientInfo)

}

func insertOneClient(client models.Clients) *mongo.InsertOneResult {
	insertResult, err := clientcollection.InsertOne(context.Background(), client)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Client Added")
	}
	return insertResult
} //insertResult contains all info

func sendEmails(lawyerEmails []string, clientInfo map[string]string) {

	fmt.Println("Goes to Emails")
	auth := smtp.PlainAuth("", "EMAIL", "PASS", "smtp.gmail.com")
	to := lawyerEmails
	msg := []byte("To: " + clientInfo["Email"] + "\r\n" +
		"Subject: Abe Consult Alert\r\n" +
		"\r\n" +
		"Colleagues, " + "\r\n" +
		"We have a customer seeking a consultation. See the details below: " +
		"Name: " + clientInfo["FirstName"] + " " + clientInfo["LastName"] + "\r\n" +
		"Location: " + clientInfo["Location"] + "\r\n" +
		"Description: " + clientInfo["Description"] + "\r\n" +
		"Are you interested in doing a consultation? If so, reply to this email ASAP.\r\n" +
		"Talib - Abe Legal Director & Co-Founder")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "EMAIL", to, msg)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Sent Email to " + strconv.Itoa(len(lawyerEmails)) + " Lawyers!")
	}
}
