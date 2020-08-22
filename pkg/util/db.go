package util

import (
	"context"
	"log"
	"time"

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
)

// ConnectDatabase to use Connect Database
func ConnectDatabase() {
	var err error
	bc := context.Background()
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
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	DB = DbConn.Database(Getenv("MONGO_DB", "shorty"))

	// options.CreateIndexes()
	// db := client.Database("shorty")
	// controllers.TodoCollection(db)
	return
}
