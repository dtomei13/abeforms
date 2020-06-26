package domain

import (
	"context"
	"fmt"
	"github.com/austinlhx/server/database"
	"github.com/austinlhx/server/models"
	"github.com/austinlhx/server/utils"
	"github.com/donvito/zoom-go/zoomAPI"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func GetCase() ([]primitive.M, *utils.ApplicationError) {
	collection := database.ClientCollection
	data, err := collection.Find(context.Background(), bson.D{{}}) //TODO: Find all with empty lawyers
	if err != nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Getting cases from DB failed"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}
	var clients []primitive.M
	for data.Next(context.Background()) {
		var lawyerEmail bson.M
		e := data.Decode(&lawyerEmail)
		if e != nil {
			log.Fatal(e)
		}
		clients = append(clients, lawyerEmail)
	}

	if err := data.Err(); err != nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Getting cases from DB failed"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}

	data.Close(context.Background())

	var openCases []primitive.M
	for _, b := range clients {
		if b[("lawyerassigned")] == "" {
			openCases = append(openCases, b)
		}
	}
	return openCases, nil
}

func GetEachCase(user models.Lawyers) ([]primitive.M, *utils.ApplicationError){
	collection := database.ClientCollection
	filter := &bson.M{"lawyerassigned": user.EmailAddress}
	data, err := collection.Find(context.TODO(), filter)
	var eachCases []primitive.M
	for data.Next(context.Background()) {
		var lawyerEmail bson.M
		e := data.Decode(&lawyerEmail)
		if e != nil {
			log.Fatal(e)
		}
		eachCases = append(eachCases, lawyerEmail)
	}
	if err != nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Getting cases from DB failed"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}
	return eachCases, nil
}

func TakeCase(caseID string, user models.Lawyers) *utils.ApplicationError{
	id, err := primitive.ObjectIDFromHex(caseID)
	if err != nil{
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Taking case failed!"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}

	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"lawyerassigned": user.EmailAddress}}
	_, error := database.ClientCollection.UpdateOne(context.Background(), filter, update)
	if error != nil{
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Updating case to DB failed!"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}
	return nil
}

func CreateMeeting(zoomInfo utils.ZoomMeeting, user models.Lawyers)*utils.ApplicationError{
	apiClient := zoomAPI.NewClient("https://api.zoom.us/v2", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6Inp6ODZTcmR0UmRLMm11TU8tTktKR0EiLCJleHAiOjE1OTI2Nzg1MDksImlhdCI6MTU5MjU5MjEwN30.AuoGUgjoI4T4YL3dnnIHjx2DS7HCp82iD-djrI4-UaE")
	userId := ("austin.abe.legal@gmail.com")
	log.Println(user.EmailAddress)
	var resp zoomAPI.CreateMeetingResponse
	var err error

	resp, err = apiClient.CreateMeeting(userId,
		"Client and Lawyer Consultation",
		2,
		"2020-06-24T22:00:00Z",
		30,
		zoomInfo.ClientEmail,
		"Asia/Singapore",
		"pass8888", //set this with your desired password for better security, max 8 chars
		"Discuss next steps and ways to contribute for this project.",
		nil,
		nil)

	if err != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Creating Zoom Meeting Failed!"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
		}
	}

	fmt.Printf("Created meeting : id = %d, topic = %s, join url = %s, start time = %s\n", resp.Id,
		resp.Topic, resp.JoinUrl, resp.StartTime)

	return nil
}