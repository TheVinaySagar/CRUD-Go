package database

import (
	"context"
	"log"
	"os"
	"time"

	"myproject/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var UserCollection *mongo.Collection
var InterviewCollection *mongo.Collection

func ConnectDB() {
	config.LoadEnv()
	MONGO_URI := os.Getenv("MONGO_URI")

	ctx, cancle := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	log.Println("Connected to MongoDB successfully")

	DB = client

	UserCollection = client.Database("InterVW").Collection("users")
	InterviewCollection = client.Database("InterVW").Collection("interviews")

}
