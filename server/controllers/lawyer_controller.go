package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"unicode"

	"../auth"
	"../models"
	"../services"
	"../utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AuthToken string

func GetLawyer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	TypeOfUser = "lawyer"

	var lawyerInfo models.Lawyers
	if err := json.NewDecoder(r.Body).Decode(&lawyerInfo); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "decoding lawyer info failed",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
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
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Email format wrong")
		return
	}
	log.Println("Email Checked")

	/*var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(lawyerInfo.Password) >= 8 {
		hasMinLen = true
	}
	for _, char := range lawyerInfo.Password {
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
	}*/

	lawyer, apiErr := services.GetLawyer(lawyerInfo.EmailAddress, lawyerInfo.Password)
	log.Print(lawyer)
	token, err := auth.GenerateJWT(*lawyer)
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
	log.Println(lawyer)
	_, error := json.Marshal(lawyer)
	if error != nil {
		apiErr := &utils.ApplicationError{
			Message:    "decoding lawyer info failed",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
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
	if err := json.NewDecoder(r.Body).Decode(&lawyerInfo); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "decoding lawyer info failed",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
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
	if format := emailFormat.MatchString(lawyerInfo.EmailAddress); format != true {
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
	if len(lawyerInfo.Password) >= 8 {
		hasMinLen = true
	}
	for _, char := range lawyerInfo.Password {
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
	if format := nameFormat.MatchString(lawyerInfo.FirstName); format != true {
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
	if format := nameFormat.MatchString(lawyerInfo.LastName); format != true {
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
	if format := phNumFormat.MatchString(lawyerInfo.PhoneNumber); format != true {
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

	//Check and validate Expertise
	if format := nameFormat.MatchString(lawyerInfo.Expertise); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "Expertise wrong format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Expertise format wrong")
		return
	}

	//Check and validate State of License
	if format := nameFormat.MatchString(lawyerInfo.StateOfLicense); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "State of License must be an URL",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("State of License format wrong")
		return
	}

	apiErr := services.AddLawyer(lawyerInfo)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
}
