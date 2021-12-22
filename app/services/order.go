package services

import (
	"context"
	"errors"
	"time"

	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOrderByID(id primitive.ObjectID) (order *models.Order, err error) {
	orderCollection := database.MI.DB.Collection("order")
	order = &models.Order{}

	query := bson.D{{Key: "id", Value: id}}
	err = orderCollection.FindOne(context.TODO(), query).Decode(order)
	return
}

func GetOrdersByUserID(userID primitive.ObjectID) (order *[]models.Order, err error) {
	orderCollection := database.MI.DB.Collection("order")
	var orders []models.Order = make([]models.Order, 0)

	query := bson.D{{Key: "userid", Value: userID}}
	cursor, err := orderCollection.Find(context.TODO(), query)
	if err != nil {
		return order, err
	}

	cursor.All(context.Background(), &orders)
	return &orders, err
}

func GetAllOrders(opts models.GetAllOrdersParams) (orders []models.Order, hasNext bool, err error) {
	orderCollection := database.MI.DB.Collection("order")

	orders = make([]models.Order, 0)
	query := bson.D{{}}

	cursor, err := orderCollection.Find(
		context.TODO(),
		query,
		options.Find().SetLimit(2),
		options.Find().SetSkip((opts.Page-1)*opts.Limit),
	)

	if err != nil {
		return orders, false, err
	}

	cursor.All(context.TODO(), &orders)
	cursor.Close(context.Background())
	remain := cursor.RemainingBatchLength()
	hasNext = true

	hasNext = true
	if remain <= 1 {
		hasNext = false
	}

	return orders, hasNext, err
}

func CreateOrder(body *models.Order, userID primitive.ObjectID) (order *models.Order, err error) {
	orderCollection := database.MI.DB.Collection("order")
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
	body.Status = models.OrderStatus.String("waiting")
	body.CreatedAt = time.Now()
	body.UpdatedAt = time.Now()

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

func ChangeOrderStatus(orderID primitive.ObjectID, orderStatus models.OrderStatus) (order *models.Order, err error) {
	orderCollection := database.MI.DB.Collection("order")

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
