package controllers

import (
	"encoding/json"
	"github.com/austinlhx/server/auth"
	"github.com/austinlhx/server/models"
	"github.com/austinlhx/server/services"
	"github.com/austinlhx/server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"regexp"
)

var AuthToken string

func GetLawyer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var lawyerInfo models.Lawyers
	if err := json.NewDecoder(r.Body).Decode(&lawyerInfo); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "decoding lawyer info failed",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil{
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Decoding lawyer info failed")
		return
	}

	emailFormat := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if format := emailFormat.MatchString(lawyerInfo.EmailAddress); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "email_address must be an email",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil{
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Email format wrong")
		return
	}
	log.Println("Email Checked")

	lawyer, apiErr := services.GetLawyer(lawyerInfo.EmailAddress, lawyerInfo.Password)
	log.Print(lawyer)
	token, err := auth.GenerateJWT(*lawyer)
	if err != nil{
		jsonValue, _ := json.Marshal(err)
		w.Write([]byte(jsonValue))
		return
	}
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	log.Println(lawyer)
	_, error := json.Marshal(lawyer)
	if error != nil{
		apiErr := &utils.ApplicationError{
			Message: "decoding lawyer info failed",
			StatusCode: http.StatusInternalServerError,
			Code: "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil{
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Something came up wrong while decoding lawyer info")
		return
	}
	AuthToken = token
}

func AddLawyer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//TODO: validating inputs

	var lawyerInfo models.Lawyers
	if err := json.NewDecoder(r.Body).Decode(&lawyerInfo); err != nil{
		apiErr := &utils.ApplicationError{
			Message: "decoding lawyer info failed",
			StatusCode: http.StatusInternalServerError,
			Code: "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil{
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Something came up wrong while decoding lawyer info")
		return
	}
	log.Println(lawyerInfo)
	lawyerInfo.ID = primitive.NewObjectID()
	log.Println(lawyerInfo.EmailAddress)

	emailFormat := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if format := emailFormat.MatchString(lawyerInfo.EmailAddress); format != true{
		apiErr := &utils.ApplicationError{
			Message: "email_address must be an email",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil{
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Something came up wrong in email")
		return
	}
	//TODO: Check all other info...
	apiErr := services.AddLawyer(lawyerInfo)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
}

