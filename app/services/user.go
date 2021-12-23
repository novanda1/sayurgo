package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/pkg/utils"
	"github.com/novanda1/sayurgo/platform/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AllUsers(c *fiber.Ctx) ([]models.User, error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)

	var users []models.User = make([]models.User, 0)
	query := bson.D{{}}

	cursor, err := userCollection.Find(c.Context(), query)
	if err != nil {
		return users, err
	}

	err = cursor.All(c.Context(), &users)

	return users, err
}

func CreateUser(data models.User) (*models.User, error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)
	data.ID = nil
	data.Role = (*string)(&models.Customer)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := userCollection.InsertOne(context.Background(), data)

	if err != nil {
		return &data, err
	}

	user := &models.User{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	userCollection.FindOne(context.TODO(), query).Decode(user)

	return user, err
}

func GetUser(userid primitive.ObjectID) (*models.User, error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)
	user := &models.User{}

	query := bson.D{{Key: "_id", Value: userid}}
	err := userCollection.FindOne(context.Background(), query).Decode(user)

	return user, err
}

func GetUserByPhone(phone string) (*models.User, error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)
	user := &models.User{}

	query := bson.D{{Key: "phone", Value: phone}}
	err := userCollection.FindOne(context.Background(), query).Decode(user)

	return user, err
}

func UpdateUser(userid primitive.ObjectID, newUserData *models.User) (*models.User, error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)

	var userDataToUpdate bson.D
	userDataToUpdate = utils.AppendOrSkip(userDataToUpdate, "displayName", newUserData.DisplayName)
	userDataToUpdate = utils.AppendOrSkip(userDataToUpdate, "phone", newUserData.Phone)
	userDataToUpdate = utils.AppendOrSkip(userDataToUpdate, "userAddress", newUserData.UserAddress)
	userDataToUpdate = utils.AppendOrSkip(userDataToUpdate, "updatedAt", time.Now())

	query := bson.D{{Key: "_id", Value: userid}}
	update := bson.D{
		{Key: "$set", Value: userDataToUpdate},
	}

	err := userCollection.FindOneAndUpdate(context.Background(), query, update).Err()
	if err != nil {
		return newUserData, err
	}

	user, _ := GetUser(userid)
	return user, err
}
