package controllers

import (
	"../auth"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"regexp"
	"unicode"

	"../models"
	"../services"
	"../utils"
)

var TypeOfUser string

func AddCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var caseInfo models.Cases
	if err := json.NewDecoder(r.Body).Decode(&caseInfo); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "error decoding claims info",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			w.WriteHeader(apiErr.StatusCode)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Error Decoding Claims Info")
		return
	}

	emailFormat := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if format := emailFormat.MatchString(caseInfo.EmailAddress); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "email_address must be an email",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Something came up wrong in email")
		return
	}

	//Check and validate first name
	nameFormat := regexp.MustCompile("^[a-zA-Z]+[a-zA-Z-']*$")
	if format := nameFormat.MatchString(caseInfo.FirstName); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "first name format is wrong",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("First Name format wrong")
		return
	}

	//Check and validate last name
	if format := nameFormat.MatchString(caseInfo.LastName); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "last name wrong",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Last Name format wrong")
		return
	}

	//Check and validate Description
	if format := nameFormat.MatchString(caseInfo.Description); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "Referral wrong format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Referral format wrong")
		return
	}

	apiErr := services.AddCase(caseInfo)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}

	emailErr := services.SendEmail()

	if emailErr != nil {
		jsonValue, err := json.Marshal(emailErr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(emailErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	log.Println("Notification sent to all lawyers!")
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	TypeOfUser = "client"
	var clientInfo models.Clients
	if err := json.NewDecoder(r.Body).Decode(&clientInfo); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "decoding client info failed",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Decoding client info failed")
		return
	}

	// CHECK AND VALIDATE EMAIL
	emailFormat := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if format := emailFormat.MatchString(clientInfo.EmailAddress); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "email_address must be an email",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Email format wrong")
		return
	}
	log.Println("Email Checked")

	//CHECK AND VALIDATE PASSWORD
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(clientInfo.Password) >= 8 {
		hasMinLen = true
	}
	for _, char := range clientInfo.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if format := hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial; format != true {
		apiErr := &utils.ApplicationError{
			Message:    "Password must be in the correct format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Password format wrong")
		return
	}

	client, apiErr := services.GetClient(clientInfo.EmailAddress, clientInfo.Password)
	log.Print(client)
	token, err := auth.GenerateClientJWT(*client)
	if err != nil {
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
	log.Println(client)
	_, error := json.Marshal(client)
	if error != nil {
		apiErr := &utils.ApplicationError{
			Message:    "decoding client info failed",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Something came up wrong while decoding client info")
		return
	}
	AuthToken = token
}

func AddClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//TODO: validating inputs

	var clientInfo models.Clients
	if err := json.NewDecoder(r.Body).Decode(&clientInfo); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "decoding client info failed",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Something came up wrong while decoding client info")
		return
	}
	log.Println(clientInfo)
	clientInfo.ID = primitive.NewObjectID()
	log.Println(clientInfo.EmailAddress)

	emailFormat := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if format := emailFormat.MatchString(clientInfo.EmailAddress); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "email_address must be an email",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Something came up wrong in email")
		return
	}
	//TODO: Check all other info...

	//CHECK AND VALIDATE PASSWORD
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(clientInfo.Password) >= 8 {
		hasMinLen = true
	}
	for _, char := range clientInfo.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if format := hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial; format != true {
		apiErr := &utils.ApplicationError{
			Message:    "Password must be in the correct format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Password format wrong")
		return
	}

	//Check and validate first name
	nameFormat := regexp.MustCompile("^[a-zA-Z]+[a-zA-Z-']*$")
	if format := nameFormat.MatchString(clientInfo.FirstName); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "first name format is wrong",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("First Name format wrong")
		return
	}

	//Check and validate last name
	if format := nameFormat.MatchString(clientInfo.LastName); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "last name wrong",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Last Name format wrong")
		return
	}

	//Check and validate phone number
	phNumFormat := regexp.MustCompile("^[\\+]?[(]?[0-9]{3}[)]?[-\\s\\.]?[0-9]{3}[-\\s\\.]?[0-9]{4,6}$")
	if format := phNumFormat.MatchString(clientInfo.PhoneNumber); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "Phone Number format wrong",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Phone Number format wrong")
		return
	}

	//Check and validate Referral/FindHow
	if format := nameFormat.MatchString(clientInfo.FindHow); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "Referral wrong format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Referral format wrong")
		return
	}

	//Check and validate Social Media
	if format := nameFormat.MatchString(clientInfo.SocialMedia); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "Social Media must be an URL",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Social Media format wrong")
		return
	}

	apiErr := services.AddClient(clientInfo)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
}
