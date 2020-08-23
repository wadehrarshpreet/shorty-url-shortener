package util

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "github.com/cavdy-play/go_mongo/controllers"
)

var (
	// DbConn connection to Mongo Client
	DbConn *mongo.Client
	// DB connection to Mongo Database
	DB *mongo.Database
	bc = context.Background()
)

// ConnectDatabase to use Connect Database
func ConnectDatabase() error {
	var err error
	// Database Config
	clientOptions := options.Client().ApplyURI(Getenv("MONGO_URI", "mongodb://localhost:27017/shorty"))
	DbConn, err = mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(bc, 10*time.Second)
	err = DbConn.Connect(ctx)
	//To close the connection at the end
	defer cancel()

	err = DbConn.Ping(bc, readpref.Primary())
	if err != nil {
		log.Println("Couldn't connect to the database")
		return err
	}
	log.Println("Connected!")

	DB = DbConn.Database(Getenv("MONGO_DB", "shorty"))

	// Ensure indexes
	indexErr := initIndexes()
	if indexErr != nil {
		log.Println("Index creation error")
		return err
	}
	return nil
}

func initIndexes() error {
	// Initialize user collection indexes
	userCollection := DB.Collection("users")
	indexName, err := userCollection.Indexes().CreateOne(bc, mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}
	log.Printf("Users Index Created %s", indexName)
	return nil
}
