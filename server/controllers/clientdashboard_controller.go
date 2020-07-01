package controllers

import (
	"../models"
	"../services"
	"../utils"
	"encoding/json"
	"log"
	"net/http"
)

func GetUnassignedCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("GetUnassignedCase Function Called")
	log.Println(r.Header)
	user := models.Clients{}
	user = r.Context().Value(utils.UserKey("user")).(models.Clients)
	unassignedCases, apiErr := services.GetUnassignedCase(user)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	json.NewEncoder(w).Encode(unassignedCases)
}

func GetAssignedCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("GetAssignedCase Function is called")
	user := models.Clients{}
	user = r.Context().Value(utils.UserKey("user")).(models.Clients)
	log.Println(user)
	log.Println("This is the user")
	assignedCase, apiErr := services.GetAssignedCase(user)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	json.NewEncoder(w).Encode(assignedCase)

}
