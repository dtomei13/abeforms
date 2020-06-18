package middleware

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://akkshay:%40Whatever12@abelegal-rrztu.gcp.mongodb.net/test"
const dbName = "AbeDB"
const collName = "clients"
const lawName = "lawyers"

var clientcollection *mongo.Collection
var lawyerCollection *mongo.Collection

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

	clientcollection = client.Database(dbName).Collection(collName)
	lawyerCollection = client.Database(dbName).Collection(lawName)

	fmt.Println("Collection instance created!")
}

func getInfo(collection *mongo.Collection) []primitive.M {
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
