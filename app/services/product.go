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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AllProducts(opts models.GetAllProductsParams) ([]models.Product, bool, error) {
	productCollection := database.MI.DB.Collection(models.ProductCollectionName)

	var products []models.Product = make([]models.Product, 0)
	query := bson.D{{}}

	cursor, err := productCollection.Find(
		context.TODO(),
		query,
		options.Find().SetLimit(opts.Limit),
		options.Find().SetSkip((opts.Page-1)*opts.Limit),
	)

	if err != nil {
		return products, false, err
	}

	cursor.All(context.TODO(), &products)
	cursor.Close(context.Background())
	remain := cursor.RemainingBatchLength()
	var hasNext bool = true

	hasNext = true
	if remain <= 1 {
		hasNext = false
	}

	return products, hasNext, err
}

func CreateProduct(body models.Product) (*models.Product, error) {
	productCollection := database.MI.DB.Collection(models.ProductCollectionName)

	body.ID = nil
	body.CreatedAt = time.Now()
	body.UpdatedAt = time.Now()

	result, err := productCollection.InsertOne(context.TODO(), body)

	if err != nil {
		return &body, err
	}

	product := &models.Product{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	productCollection.FindOne(context.TODO(), query).Decode(product)

	return product, err
}

func GetProduct(id primitive.ObjectID) (*models.Product, error) {
	productCollection := database.MI.DB.Collection(models.ProductCollectionName)
	product := &models.Product{}

	query := bson.D{{Key: "_id", Value: id}}
	err := productCollection.FindOne(context.TODO(), query).Decode(product)

	return product, err
}

func UpdateProduct(id primitive.ObjectID, data *models.Product) (*models.Product, error) {
	productCollection := database.MI.DB.Collection(models.ProductCollectionName)

	query := bson.D{{Key: "_id", Value: id}}

	// store the data that need to update
	var dataToUpdate bson.D

	dataToUpdate = utils.AppendOrSkip(dataToUpdate, "title", data.Title)
	dataToUpdate = utils.AppendOrSkip(dataToUpdate, "imageUrl", data.ImageUrl)
	dataToUpdate = utils.AppendOrSkip(dataToUpdate, "price", data.Price)
	dataToUpdate = utils.AppendOrSkip(dataToUpdate, "discountPrice", data.DiscountPrice)
	dataToUpdate = utils.AppendOrSkip(dataToUpdate, "categories", data.Categories)
	dataToUpdate = utils.AppendOrSkip(dataToUpdate, "unitType", data.UnitType)
	dataToUpdate = utils.AppendOrSkip(dataToUpdate, "information", data.Information)
	dataToUpdate = utils.AppendOrSkip(dataToUpdate, "nutrition", data.Nutrition)
	dataToUpdate = utils.AppendOrSkip(dataToUpdate, "updatedAt", time.Now())

	update := bson.D{
		{Key: "$set", Value: dataToUpdate},
	}

	err := productCollection.FindOneAndUpdate(context.TODO(), query, update).Err()
	if err != nil {
		return data, err
	}

	product, err := GetProduct(id)
	return product, err
}

func DeleteProduct(c *fiber.Ctx) error {
	productCollection := database.MI.DB.Collection(models.ProductCollectionName)

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
