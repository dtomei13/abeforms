package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strconv"

	"../models"

	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://akkshay:%40Whatever12@abelegal-rrztu.gcp.mongodb.net/test"
const dbName = "AbeDB"
const collName = "clients"
const lawName = "lawyers"

var collection *mongo.Collection
var lawyerCollection *mongo.Collection
var key = make([]byte, 64)
var store = sessions.NewCookieStore(key)

func init() {

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database is up and running")

	collection = client.Database(dbName).Collection(collName)
	lawyerCollection = client.Database(dbName).Collection(lawName)

	fmt.Println("Collection instance created!")
}

//Template for future projects -> Grabs all clients infos
func getClientsInfo() []primitive.M {

	data, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var clients []primitive.M
	for data.Next(context.Background()) {
		var client bson.M
		e := data.Decode(&client)
		if e != nil {
			log.Fatal(e)
		}
		clients = append(clients, client)

	}
	if err := data.Err(); err != nil {
		log.Fatal(err)
	}

	data.Close(context.Background())
	return clients
}

// GetAllLawyerEmails Gets all the email addresses of lawyers from the database
func GetAllLawyerEmails() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func sendEmails(lawyerEmails []string, clientInfo map[string]string) {

	fmt.Println("Goes to Emails")
	auth := smtp.PlainAuth("", "EMAIL", "PASS", "smtp.gmail.com")
	to := lawyerEmails
	msg := []byte("To: " + clientInfo["Email"] + "\r\n" +
		"Subject: Abe Consult Alert\r\n" +
		"\r\n" +
		"Colleagues, " + "\r\n" +
		"We have a customer seeking a consultation. See the details below: " +
		"Name: " + clientInfo["FirstName"] + " " + clientInfo["LastName"] + "\r\n" +
		"Location: " + clientInfo["Location"] + "\r\n" +
		"Description: " + clientInfo["Description"] + "\r\n" +
		"Are you interested in doing a consultation? If so, reply to this email ASAP.\r\n" +
		"Talib - Abe Legal Director & Co-Founder")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "EMAIL", to, msg)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Sent Email to " + strconv.Itoa(len(lawyerEmails)) + " Lawyers!")
	}
}

// CreateClientsInfo Posts info of client to the next task at hand
func CreateClientsInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var client models.Clients
	_ = json.NewDecoder(r.Body).Decode(&client)
	insertOneClient(client)

	json.NewEncoder(w).Encode(client)
	lawyersPrimitive := getAllLawyers()
	var lawyersEmails []string

	for _, b := range lawyersPrimitive {
		lawyersEmails = append(lawyersEmails, b[("email")].(string))
	}

	clientInfo := map[string]string{
		"FirstName":    client.FirstName,
		"LastName":     client.LastName,
		"EmailAddress": client.EmailAddress,
		"PhoneNumber":  client.PhoneNumber,
		"Location":     client.StateOfIssue,
		"Description":  client.Description,
		"FindHow":      client.FindHow,
		"SocialMedia":  client.SocialMedia,
	}

	fmt.Println(lawyersEmails)
	fmt.Println(clientInfo)

	sendEmails(lawyersEmails, clientInfo)

}

func insertOneClient(client models.Clients) *mongo.InsertOneResult {
	insertResult, err := collection.InsertOne(context.Background(), client)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Client Added")
	}
	return insertResult
} //insertResult contains all info

// get all task from the DB and return it
func getAllLawyers() []primitive.M {
	cur, err := lawyerCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M //bson Map object {"a": 1, "b": true}
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// results will contain the emails of lawyers
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results //primitive.M object of bSON objects
} //insertResult contains all info

// get all task from the DB and return it

func insertOneLawyer(lawyer models.LawyerSignUp) *mongo.InsertOneResult {
	insertResult, err := lawyerCollection.InsertOne(context.Background(), lawyer)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Client Added")
	}
	return insertResult
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

/*
	lawyersPrimitive := getAllLawyers()

	lawyerInfo := map[string]string{
		"FirstName":      lawyerSignUp.FirstName,
		"LastName":       lawyerSignUp.LastName,
		"EmailAddress":   lawyerSignUp.EmailAddress,
		"PhoneNumber":    lawyerSignUp.PhoneNumber,
		"StateOfLicense": lawyerSignUp.StateOfLicense,
		"Expertise":      lawyerSignUp.Expertise,
		"Password":       lawyerSignUp.Password,
		"RetypePassword": lawyerSignUp.RetypePassword,
	}
	fmt.Println(lawyerInfo)
	fmt.Println(lawyersPrimitive)*/

func LawyerSignIn(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	session, _ := store.Get(r, "cookie-name")

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

func SendAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(true)
	fmt.Println("sent")
}

func lawyerAuth(email string, pass string, w http.ResponseWriter, r *http.Request) {
	lawyersPrimitive := getAllLawyers()
	session, _ := store.Get(r, "cookie-name")

	for _, b := range lawyersPrimitive {
		if email == b[("emailaddress")].(string) && pass == b[("password")].(string) {
			fmt.Println("Signed In")
			session.Values["authenticated"] = true
			session.Save(r, w)
		}
	}
}
