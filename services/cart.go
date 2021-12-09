package services

import (
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

func AddProductToCart(c *fiber.Ctx) (*models.Cart, error) {
	productCollection := config.MI.DB.Collection("carts")
	data := new(models.Cart)
	cartProduct := new(models.CartProduct)

	paramId := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return data, err
	}

	err = c.BodyParser(&cartProduct)
	if err != nil {
		return data, err
	}

	query := bson.M{"_id": id, "product": bson.M{"$in": cartProduct.ID}}

	var newCartProduct bson.D

	if cartProduct.ProductID != nil && cartProduct.TotalProduct != nil {
		newCartProduct = append(newCartProduct, bson.E{Key: "productID", Value: cartProduct.ProductID})
		newCartProduct = append(newCartProduct, bson.E{Key: "totalProduct", Value: cartProduct.TotalProduct})
	}

	update := bson.D{
		{Key: "$push", Value: bson.M{"product": newCartProduct}},
	}

	err = productCollection.FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		return data, err
	}

	cart, err := GetCart(c, paramId)
	return cart, err
}
