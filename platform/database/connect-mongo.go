package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/novanda1/sayurgo/app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DATABASE_NAME")),
	}

	CreateMongoIndex(models.UserCollectionName, "phone", true)
	CreateMongoIndex(models.CartCollectionName, "userid", true)
}
