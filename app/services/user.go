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

func GetUser(c *fiber.Ctx, paramId string) (*models.User, error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)
	user := &models.User{}

	id, err := primitive.ObjectIDFromHex(paramId)

	if err != nil {
		return user, err
	}

	query := bson.D{{Key: "_id", Value: id}}
	err = userCollection.FindOne(c.Context(), query).Decode(user)

	return user, err
}

func GetUserByPhone(phone string) (*models.User, error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)
	user := &models.User{}

	query := bson.D{{Key: "phone", Value: phone}}
	err := userCollection.FindOne(context.Background(), query).Decode(user)

	return user, err
}

func UpdateUser(c *fiber.Ctx) (*models.User, error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)
	data := new(models.User)

	paramId := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return data, err
	}

	err = c.BodyParser(&data)
	if err != nil {
		return data, err
	}

	var userDataToUpdate bson.D
	userDataToUpdate = utils.AppendOrSkip(userDataToUpdate, "displayName", data.DisplayName)
	userDataToUpdate = utils.AppendOrSkip(userDataToUpdate, "phone", data.Phone)
	userDataToUpdate = utils.AppendOrSkip(userDataToUpdate, "userAddress", data.UserAddress)
	userDataToUpdate = utils.AppendOrSkip(userDataToUpdate, "updatedAt", time.Now())

	query := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: userDataToUpdate},
	}

	err = userCollection.FindOneAndUpdate(c.Context(), query, update).Err()
	if err != nil {
		return data, err
	}

	user, _ := GetUser(c, paramId)
	return user, err

}
