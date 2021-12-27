package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/pkg/utils"
	"github.com/novanda1/sayurgo/platform/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func IsAddressExists(userid primitive.ObjectID, title string) (exists bool, err error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)

	query := bson.M{"userAddress.title": title}

	var (
		opts        options.FindOptions
		modelStruct []models.UserAddress
	)

	opts.SetProjection(bson.M{"userAddress": "1"})

	cursor, err := userCollection.Find(context.Background(), query, &opts)
	if err != nil {
		return false, err
	}

	cursor.All(context.Background(), &modelStruct)

	if modelStruct != nil {
		return true, err
	} else {
		return false, err
	}

}

func AddUserAddress(userId primitive.ObjectID, body models.UserAddress) (user *models.User, err error) {
	userCollection := database.MI.DB.Collection(models.UserCollectionName)

	addressTitleExists, _ := IsAddressExists(userId, *body.Title)
	if addressTitleExists {
		return user, errors.New("address is already exists")
	}

	user, err = GetUser(userId)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	query := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{}

	var addresses = user.UserAddress
	if addresses == nil {
		var addresses []models.UserAddress
		addresses = append(addresses, body)
		update = append(update, bson.E{Key: "$set", Value: bson.M{"userAddress": addresses}})
	} else if len(*addresses) >= 3 {
		err = errors.New("address can't more than 3")
		return user, err
	} else {
		update = append(update, bson.E{Key: "$push", Value: bson.M{"userAddress": body}})
	}

	err = userCollection.FindOneAndUpdate(context.Background(), query, update).Err()
	if err != nil {
		return user, err
	}

	user, _ = GetUser(userId)
	return user, err
}
