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

func AllProducts() (error, []models.Product) {
	productCollection := config.MI.DB.Collection("products")

	var products []models.Product = make([]models.Product, 0)
	query := bson.D{{}}

	cursor, err := productCollection.Find(context.TODO(), query)

	if err != nil {
		return err, products
	}

	cursor.All(context.TODO(), &products)

	return err, products
}

func CreateProduct(body models.Product) (error, *models.Product) {
	productCollection := config.MI.DB.Collection("products")

	body.ID = nil
	body.CreatedAt = time.Now()
	body.UpdatedAt = time.Now()

	result, err := productCollection.InsertOne(context.TODO(), body)

	if err != nil {
		return err, &body
	}

	product := &models.Product{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	productCollection.FindOne(context.TODO(), query).Decode(product)

	return err, product
}

func GetProduct(id primitive.ObjectID) (error, *models.Product) {
	productCollection := config.MI.DB.Collection("products")
	product := &models.Product{}

	query := bson.D{{Key: "_id", Value: id}}
	err := productCollection.FindOne(context.TODO(), query).Decode(product)

	return err, product
}

func UpdateProduct(id primitive.ObjectID, data *models.Product) (error, *models.Product) {
	productCollection := config.MI.DB.Collection("products")

	query := bson.D{{Key: "_id", Value: id}}

	// store the data that need to update
	var dataToUpdate bson.D

	if data.Title != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "title", Value: data.Title})
	}

	if data.ImageUrl != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "imageUrl", Value: data.ImageUrl})
	}

	if data.Price != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "price", Value: data.Price})
	}

	if data.Categories != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "discountPrice", Value: data.DiscountPrice})
	}

	if data.Categories != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "categories", Value: data.Categories})
	}

	if data.UnitType != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "unitType", Value: data.UnitType})
	}

	if data.Nutrition != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "information", Value: data.Information})
	}

	if data.Nutrition != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "nutrition", Value: data.Nutrition})
	}

	dataToUpdate = append(dataToUpdate, bson.E{Key: "updatedAt", Value: time.Now()})

	update := bson.D{
		{Key: "$set", Value: dataToUpdate},
	}

	err := productCollection.FindOneAndUpdate(context.TODO(), query, update).Err()
	if err != nil {
		return err, data
	}

	_, product := GetProduct(id)
	return err, product
}

func DeleteProduct(c *fiber.Ctx) error {
	productCollection := config.MI.DB.Collection("products")

	paramID := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}

	err = productCollection.FindOneAndDelete(c.Context(), query).Err()

	return err
}
