package services

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCart(cart models.Cart) (*models.Cart, error) {
	cartCollection := config.MI.DB.Collection("carts")

	cart.ID = nil
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()

	result, err := cartCollection.InsertOne(context.TODO(), cart)

	if err != nil {
		return &cart, err
	}

	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	cartCollection.FindOne(context.TODO(), query).Decode(cart)

	return &cart, err
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

func AddProductToCart(c *fiber.Ctx, userID string, cartProduct *models.CartProduct) (cart *models.Cart, message string) {
	cartCollection := config.MI.DB.Collection("carts")
	query := bson.M{"userid": userID}

	var update bson.D

	if cart.Product != nil {
		update = append(update, bson.E{"$push", bson.M{"product": cartProduct.ID}})
	} else {
		var cartProductArray [1]string
		cartProductArray[0] = *cartProduct.ID
		update = append(update, bson.E{"$set", bson.M{"product": cartProductArray}})

	}

	err := cartCollection.FindOne(context.TODO(), query).Decode(&cart)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			cartOption := &models.Cart{}
			cartOption.UserID = &userID

			cart, err = CreateCart(*cartOption)

			if err != nil {
				return cart, "failed to create cart"
			}
		}
	}

	err = cartCollection.FindOneAndUpdate(context.Background(), query, update).Decode(cart)
	fmt.Println(err, cart)

	if err != nil {
		return cart, "update failed"
	}

	return cart, "successfully"

}
