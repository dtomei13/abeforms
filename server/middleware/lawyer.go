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
var lawyerID primitive.ObjectID

func CaseComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	caseComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
} //LawyerAssigned

func caseComplete(lawCase string) {
	fmt.Println(lawCase)
	id, _ := primitive.ObjectIDFromHex(lawCase)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"lawyerassigned": lawyerEmail}}
	result, err := clientcollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func GetOpenCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	payload := getOpenCases()
	json.NewEncoder(w).Encode(payload)

}

func getOpenCases() []primitive.M {

	clientsPrimitive := getInfo(clientcollection)
	var clients []primitive.M
	for _, b := range clientsPrimitive {
		if b[("lawyerassigned")] == "" {
			clients = append(clients, b)
		}
	}
	return clients
}

func GetMyCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	payload := getMyCases()
	json.NewEncoder(w).Encode(payload)

}

func getMyCases() []primitive.M {

	clientsPrimitive := getInfo(clientcollection)
	var clients []primitive.M
	for _, b := range clientsPrimitive {
		if b[("lawyerassigned")] == lawyerEmail {
			clients = append(clients, b)
		}
	}
	return clients
}

func insertOneLawyer(lawyer models.LawyerSignUp) *mongo.InsertOneResult {
	//TODO: What if they already exist?
	//TODO: Check if lawyer is existing user
	//TODO: Retrieve last lawyerid, to +1
	insertResult, err := lawyerCollection.InsertOne(context.Background(), lawyer)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Client Added")
	}
	return insertResult
}

//TODO: Sanitize on a security side
//TODO: Sanitize All user Input -> Validating email address, and info on both ends DO NOT TRUST ANYTHING FROM FRONTEND, Does lawyer exist?

func LawyerSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var lawyerSignUp models.LawyerSignUp
	_ = json.NewDecoder(r.Body).Decode(&lawyerSignUp)
	lawyerSignUp.ID = primitive.NewObjectID()
	fmt.Println(lawyerSignUp)
	insertOneLawyer(lawyerSignUp)

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
	for _, b := range lawyersPrimitive { //Brute force -> TODO: Change in the future
		if lawyerSignIn.EmailAddress == b[("emailaddress")].(string) && lawyerSignIn.Password == b[("password")].(string) { //Plain Text password... Security Issue
			fmt.Println("Signed In")
			//lawyerEmail = lawyerSignIn.EmailAddress
			//lawyerPassword = lawyerSignIn.Password
			lawyerEmail = lawyerSignIn.EmailAddress
			lawyerID = lawyerSignIn.ID
			fmt.Println(lawyerEmail)
			fmt.Println(lawyerID)
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

func SendAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	token, err := CreateToken(lawyerID, lawyerEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("StatusOK", token)

	payload := map[string]interface{}{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"user":          token.EmailAddress,
		"expiry":        token.AtExpires,
	}
	fmt.Println(payload["expiry"])
	json.NewEncoder(w).Encode(payload)
	fmt.Println("sent")

}

func CreateToken(userid primitive.ObjectID, email string) (*models.TokenDetails, error) {

	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.EmailAddress = email

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //Set environment TODO: Generate a random 64byte key
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //Set environment TODO: Generate a random 64byte key
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	return td, nil
}
