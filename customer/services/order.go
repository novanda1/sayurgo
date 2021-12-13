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
	orderCollection := config.MI.DB.Collection("order")
	order = new(models.Order)

	cart, err := GetCart(userID)
	if err != nil {
		return body, err
	}

	var totalPrice int
	var orderProducts = []models.OrderProduct{}
	for _, c := range *cart.Product {
		product, _ := GetProduct(c.ProductID)
		var price int

		if product.DiscountPrice != nil {
			price = *product.DiscountPrice
		} else {
			price = *product.Price
		}

		var orderProduct models.OrderProduct

		orderProduct.ProductID = c.ProductID
		orderProduct.TotalProduct = c.TotalProduct
		orderProduct.AtPrice = &price

		totalPrice += price
		orderProducts = append(orderProducts, orderProduct)
	}

	body.Products = &orderProducts
	body.UserID = userID
	body.TotalPrice = &totalPrice
	result, err := orderCollection.InsertOne(context.TODO(), body)

	if err != nil {
		return body, err
	}

	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	err = orderCollection.FindOne(context.Background(), query).Decode(order)

	if err == nil {
		ClearProductsInCart(userID)
	}

	return order, err
}
