package database

import "go.mongodb.org/mongo-driver/mongo"

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance
