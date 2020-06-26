package controllers

import (
	"encoding/json"
	"github.com/austinlhx/server/models"
	"github.com/austinlhx/server/services"
	"github.com/austinlhx/server/utils"
	"log"
	"net/http"
	"regexp"
)

func AddClient(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var clientInfo models.Clients
	if err := json.NewDecoder(r.Body).Decode(&clientInfo); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "error decoding client info",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil{
			w.WriteHeader(apiErr.StatusCode)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Error Decoding Client Info")
		return
	}

	emailFormat := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if format := emailFormat.MatchString(clientInfo.EmailAddress); format != true {
		apiErr := &utils.ApplicationError{
			Message:    "email_address must be an email",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil{
			log.Println("Error decoding json")
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Email format wrong")
		return
	}
	log.Println("Email Checked")

	apiErr := services.AddClient(clientInfo)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}

	emailErr := services.SendEmail()

	if emailErr != nil {
		jsonValue, err := json.Marshal(emailErr)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(emailErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	log.Println("Notification sent to all lawyers!")
}