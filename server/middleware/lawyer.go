package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertOneLawyer(lawyer models.LawyerSignUp) *mongo.InsertOneResult {
	insertResult, err := lawyerCollection.InsertOne(context.Background(), lawyer)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Client Added")
	}
	return insertResult
}

func CaseComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	session, _ := store.Get(r, "loginSession")
	params := mux.Vars(r)
	fmt.Println(session.Values)
	caseComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func caseComplete(lawCase string) {
	fmt.Println(lawCase)
	id, _ := primitive.ObjectIDFromHex(lawCase)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := clientcollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func LawyerSignIn(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	session, _ := store.Get(r, "loginSession")

	var lawyerSignIn models.LawyerSignIn
	_ = json.NewDecoder(r.Body).Decode(&lawyerSignIn)
	lawyerAuth(lawyerSignIn.EmailAddress, lawyerSignIn.Password, w, r)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		fmt.Println("ERROR")
		return
	}
	fmt.Println("YAY")

	SendAuth(w, r)

}

func GetCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	payload := getInfo(clientcollection)

	json.NewEncoder(w).Encode(payload)

}

func lawyerAuth(email string, pass string, w http.ResponseWriter, r *http.Request) {
	lawyersPrimitive := getInfo(lawyerCollection)
	session, err := store.Get(r, "loginSession")

	for _, b := range lawyersPrimitive {
		if email == b[("emailaddress")].(string) && pass == b[("password")].(string) {
			fmt.Println("Signed In")
			session.Values["authenticated"] = true
			session.Values["EmailAddress"] = email
			break
		}
	}
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LawyerSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var lawyerSignUp models.LawyerSignUp
	_ = json.NewDecoder(r.Body).Decode(&lawyerSignUp)
	insertOneLawyer(lawyerSignUp)

	json.NewEncoder(w).Encode(lawyerSignUp)

}

func SendAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(true)
	fmt.Println("sent")
}
