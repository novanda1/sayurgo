package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AllProducts(c *fiber.Ctx) (error, []models.Product) {
	productCollection := config.MI.DB.Collection("products")

	var products []models.Product = make([]models.Product, 0)
	query := bson.D{{}}

	cursor, err := productCollection.Find(c.Context(), query)

	if err != nil {
		return err, products
	}

	cursor.All(c.Context(), &products)

	return err, products
}

func CreateProduct(c *fiber.Ctx) (error, *models.Product) {
	productCollection := config.MI.DB.Collection("products")
	data := new(models.Product)

	err := c.BodyParser(&data)
	if err != nil {
		return err, data
	}

	data.ID = nil
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := productCollection.InsertOne(c.Context(), data)

	if err != nil {
		return err, data
	}

	product := &models.Product{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	productCollection.FindOne(c.Context(), query).Decode(product)

	return err, product
}

func GetProduct(c *fiber.Ctx, paramId string) (error, *models.Product) {
	productCollection := config.MI.DB.Collection("products")
	product := &models.Product{}

	id, err := primitive.ObjectIDFromHex(paramId)

	if err != nil {
		return err, product
	}

	query := bson.D{{Key: "_id", Value: id}}
	err = productCollection.FindOne(c.Context(), query).Decode(product)

	return err, product
}

func UpdateProduct(c *fiber.Ctx) (error, *models.Product) {
	productCollection := config.MI.DB.Collection("products")
	data := new(models.Product)

	paramId := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return err, data
	}

	err = c.BodyParser(&data)
	if err != nil {
		return err, data
	}

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

	err = productCollection.FindOneAndUpdate(c.Context(), query, update).Err()
	if err != nil {
		return err, data
	}

	_, product := GetProduct(c, paramId)
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
