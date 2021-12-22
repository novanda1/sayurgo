package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoIndex(collectionName string, field string, unique bool) bool {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := MI.DB.Collection(collectionName)

	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
