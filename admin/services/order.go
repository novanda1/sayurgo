package adminServices

import (
	"context"
	"errors"

	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChangeOrderStatus(orderID primitive.ObjectID, orderStatus models.OrderStatus) (order *models.Order, err error) {
	orderCollection := config.MI.DB.Collection("order")

	query := bson.M{"_id": orderID}
	err = orderCollection.FindOne(context.Background(), query).Decode(&order)
	if err != nil {
		return
	}

	if order.Status == orderStatus.String() {
		return order, errors.New("order status is already '" + order.Status + "' and you assign me to change it to '" + orderStatus.String() + "'")
	}

	query = bson.M{"_id": orderID}
	update := bson.M{"$set": bson.M{"orderStatus": orderStatus.String()}}
	_, err = orderCollection.UpdateOne(context.Background(), query, update)
	if err != nil {
		return
	}

	query = bson.M{"_id": orderID}
	err = orderCollection.FindOne(context.Background(), query).Decode(&order)
	if err != nil {
		return
	}

	return
}
