package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const connectionString = "mongodb+srv://akkshay:%40Whatever12@abelegal-rrztu.gcp.mongodb.net/test"
const dbName = "AbeDB"
const collName = "client_info"
const lawName = "lawyers"
const claimName = "clients"

var ClientCollection *mongo.Collection
var LawyerCollection *mongo.Collection
var ClaimsCollection *mongo.Collection

func ConnectDB() {

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

	ClientCollection = client.Database(dbName).Collection(collName)
	LawyerCollection = client.Database(dbName).Collection(lawName)
	ClaimsCollection = client.Database(dbName).Collection(claimName)

	fmt.Println("Collection instance created!")

}
