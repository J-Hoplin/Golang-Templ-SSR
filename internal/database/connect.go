package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
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
	dburi = os.Getenv("DATABASE_URL")
	dbname = os.Getenv("DATABASE_NAME")

	// Check env variable status
	if dburi == "" {
		log.Fatalln("Database URI Not found")
	} else {
		if dbname == "" {
			log.Fatalln("Database Name not found")
		}
	}

	if connection, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi)); err != nil {
		log.Fatalln("Fail to connect database:", err.Error())
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
