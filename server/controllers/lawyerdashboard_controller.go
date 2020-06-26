package controllers

import (
	"encoding/json"
	"github.com/austinlhx/server/models"
	"github.com/austinlhx/server/services"
	"github.com/austinlhx/server/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetCase(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("GetCase Function Called")
	log.Println(r.Header)
	openCases, apiErr := services.GetCase()
	if apiErr != nil{
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	json.NewEncoder(w).Encode(openCases)
}

func GetEachCase(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("GetEachCase Function is called")
	user := models.Lawyers{}
	user = r.Context().Value(utils.UserKey("user")).(models.Lawyers)
	log.Println(user)
	log.Println("This is the user")
	eachCase, apiErr := services.GetEachCase(user)
	if apiErr != nil{
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	json.NewEncoder(w).Encode(eachCase)

}

func CreateMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	user := models.Lawyers{}
	user = r.Context().Value(utils.UserKey("user")).(models.Lawyers)

	var zoomMeetingInfo utils.ZoomMeeting
	if err := json.NewDecoder(r.Body).Decode(&zoomMeetingInfo); err != nil{
		apiErr := &utils.ApplicationError{
			Message: "decoding zoom_meeting info failed",
			StatusCode: http.StatusInternalServerError,
			Code: "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil{
			log.Println(err)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Something came up wrong decoding zoom meeting info")
		return
	}

	apiErr := services.CreateMeeting(zoomMeetingInfo, user)

	if apiErr != nil{
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}

}

func TakeCase(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params:= mux.Vars(r)
	user := models.Lawyers{}
	user = r.Context().Value(utils.UserKey("user")).(models.Lawyers)
	apiErr := services.TakeCase(params["id"], user)
	if apiErr != nil{
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	log.Printf("Case Taken by %v", user.EmailAddress)

}