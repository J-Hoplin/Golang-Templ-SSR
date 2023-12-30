package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collections

type DBCollection string

const (
	TODO DBCollection = "todo"
)

// Client information
var client *mongo.Client
var dburi, dbname string

func Connect() {
	dburi = os.Getenv("MONGODB_URI")
	dbname = os.Getenv("MONGODB_DATABASE_NAME")

	// Check env variable status
	if dburi == "" {
		panic("Database URI Not found")
	} else {
		if dbname == "" {
			panic("Database Name not found")
		}
	}

	if connection, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi)); err != nil {
		panic(fmt.Sprintln("Fail to connect database:", err.Error()))
	} else {
		client = connection
	}
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Println("Fail to disconnect database:", err.Error())
	}
}

// Get client, database, collection
func GetClient() *mongo.Client {
	return client
}

func GetDatabase() *mongo.Database {
	return client.Database(dbname)
}

func GetCollection(collection DBCollection) *mongo.Collection {
	return client.Database(dbname).Collection(string(collection))
}
