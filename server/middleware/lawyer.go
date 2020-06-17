package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"../models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var lawyerEmail string

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
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
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

func GetCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	payload := getInfo(clientcollection)

	json.NewEncoder(w).Encode(payload)

}

func SendAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	token, err := CreateToken(lawyerEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("StatusOK", token)

	payload := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"user":          token.EmailAddress,
	}
	json.NewEncoder(w).Encode(payload)
	fmt.Println("sent")
}

func LawyerSignIn(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var lawyerSignIn models.LawyerSignIn
	_ = json.NewDecoder(r.Body).Decode(&lawyerSignIn)

	lawyersPrimitive := getInfo(lawyerCollection)

	var authen bool
	authen = false
	//var lawyerPassword string
	//var lawyerID primitive.ObjectID
	for _, b := range lawyersPrimitive {
		if lawyerSignIn.EmailAddress == b[("emailaddress")].(string) && lawyerSignIn.Password == b[("password")].(string) {
			fmt.Println("Signed In")
			//lawyerEmail = lawyerSignIn.EmailAddress
			//lawyerPassword = lawyerSignIn.Password
			lawyerEmail = lawyerSignIn.EmailAddress
			fmt.Println(lawyerEmail)
			authen = true

			break
		}
	}
	if authen == false {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
	//no Lawyer found
	fmt.Println(lawyerEmail)
}

func CreateToken(userid string) (*models.TokenDetails, error) {
	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.EmailAddress = userid

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}
