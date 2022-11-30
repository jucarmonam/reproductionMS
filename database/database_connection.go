package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 20
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

var mc *mongo.Client
var database *mongo.Database

//DBinstance func
func DBinstance() *mongo.Database{
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("REPRODUCTION_DB_USER")
	password := os.Getenv("REPRODUCTION_DB_PASSWORD")
	clusterEndpoint := os.Getenv("REPRODUCTION_DB_ENDPOINT")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	defer cancel()

	var err error
	mc, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
	}
	database = mc.Database("crowstream_reproduction_db")

	fmt.Println("Connected to MongoDB!")

	return database
}

//Client Database instance
var Client = DBinstance()


//OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Database, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Collection(collectionName)

	return collection
}
