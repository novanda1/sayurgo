package services

import (
	"context"

	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderByID(id primitive.ObjectID) (order *models.Order, err error) {
	orderCollection := config.MI.DB.Collection("order")
	order = &models.Order{}

	query := bson.D{{Key: "id", Value: id}}
	err = orderCollection.FindOne(context.TODO(), query).Decode(order)
	return
}

func CreateOrder(body *models.Order, userID primitive.ObjectID) (order *models.Order, err error) {
	// productCollection := config.MI.DB.Collection("product")
	orderCollection := config.MI.DB.Collection("order")
	order = new(models.Order)

	cart, err := GetCart(userID)
	if err != nil {
		return body, err
	}

	var orderProducts = []models.OrderProduct{}
	for _, c := range *cart.Product {
		product, _ := GetProduct(c.ProductID)

		var orderProduct models.OrderProduct

		orderProduct.ProductID = c.ProductID
		orderProduct.TotalProduct = c.TotalProduct
		orderProduct.AtPrice = product.Price
		orderProducts = append(orderProducts, orderProduct)
	}

	body.Products = &orderProducts
	result, err := orderCollection.InsertOne(context.TODO(), body)

	if err != nil {
		return body, err
	}

	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	err = orderCollection.FindOne(context.Background(), query).Decode(order)
	return order, err
}
