package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCart(c *fiber.Ctx) (*models.Cart, error) {
	cartCollection := config.MI.DB.Collection("carts")
	data := new(models.Cart)

	err := c.BodyParser(&data)
	if err != nil {
		return data, err
	}

	data.ID = nil
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := cartCollection.InsertOne(c.Context(), data)

	if err != nil {
		return data, err
	}

	cart := &models.Cart{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	cartCollection.FindOne(c.Context(), query).Decode(cart)

	return cart, err
}

func GetCart(c *fiber.Ctx, paramId string) (*models.Cart, error) {
	cartCollection := config.MI.DB.Collection("carts")
	cart := &models.Cart{}

	id, err := primitive.ObjectIDFromHex(paramId)

	if err != nil {
		return cart, err
	}

	query := bson.D{{Key: "_id", Value: id}}
	err = cartCollection.FindOne(c.Context(), query).Decode(cart)

	return cart, err
}

func AddProductToCart(user *models.User, cartProduct *models.CartProduct) (cart *models.Cart, message string) {
	cartCollection := config.MI.DB.Collection("carts")
	query := bson.M{"userid": user.ID}

	update := bson.D{
		{Key: "$push", Value: bson.M{"product": cartProduct}},
	}

	err := cartCollection.FindOneAndUpdate(context.TODO(), query, update).Decode(cart)

	if err != nil {
		return cart, "failed to update"
	}

	return cart, "successfully"

}
