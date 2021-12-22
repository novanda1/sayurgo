package services

import (
	"context"
	"time"

	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/platforms/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCart(cart models.Cart) (*models.Cart, error) {
	cartCollection := database.MI.DB.Collection("carts")

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

func GetCart(userID primitive.ObjectID) (*models.Cart, error) {
	cartCollection := database.MI.DB.Collection("carts")
	cart := &models.Cart{}

	query := bson.D{{Key: "userid", Value: userID}}
	err := cartCollection.FindOne(context.TODO(), query).Decode(cart)

	return cart, err
}

func AddProductToCart(userID primitive.ObjectID, cartProduct *models.CartProduct) (cart *models.Cart, message string) {
	cartCollection := database.MI.DB.Collection("carts")
	query := bson.M{"userid": userID}

	cart, err := GetCart(userID)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			cartOption := &models.Cart{}
			cartOption.UserID = userID

			cart, err = CreateCart(*cartOption)

			if err != nil {
				return cart, "failed to create cart"
			}
		}
	}

	var dataToUpdate bson.D
	dataToUpdate = append(dataToUpdate, bson.E{Key: "product", Value: cartProduct})

	var update bson.D
	update = append(update, primitive.E{Key: "$push", Value: dataToUpdate})

	err = cartCollection.FindOneAndUpdate(context.TODO(), query, update).Decode(&models.Cart{})
	if err != nil {
		return cart, "update failed"
	}

	cart, _ = GetCart(userID)

	return cart, "successfully"

}

func IsProductAlreadyExists(productID primitive.ObjectID, userID primitive.ObjectID) bool {
	cartCollection := database.MI.DB.Collection("carts")
	_, err := GetCart(userID)
	if err != nil {
		return false
	}

	count, err := cartCollection.CountDocuments(context.Background(), bson.M{"userid": userID, "product.productID": productID})
	if err != nil {
		return false
	}

	if count != 0 {
		return true
	}

	return false
}

func DeleteProductFromCart(productID primitive.ObjectID, userID primitive.ObjectID) (message string, success bool) {
	cartCollection := database.MI.DB.Collection("carts")
	isExists := IsProductAlreadyExists(productID, userID)
	if !isExists {
		message = "product didn't even exists"
		success = false
		return
	}

	query := bson.M{"userid": userID}
	update := bson.M{"$pull": bson.M{"product": bson.M{"productID": productID}}}

	_, err := cartCollection.UpdateMany(context.TODO(), query, update)

	if err != nil {
		message = "something went wrong when update"
		success = false
		return
	}

	message = "product deleted successfully"
	success = true
	return
}

func ChangeTotalProductInCart(productID primitive.ObjectID, userID primitive.ObjectID, totalProduct int) (message string, success bool, data *models.Cart) {
	cartCollection := database.MI.DB.Collection("carts")
	isExists := IsProductAlreadyExists(productID, userID)
	if !isExists {
		message = "product didn't even exists"
		success = false
		return
	}

	query := bson.M{"userid": userID, "product.productID": productID}
	update := bson.M{"$set": bson.M{"product.$.totalProduct": totalProduct}}

	_, err := cartCollection.UpdateOne(context.TODO(), query, update)

	if err != nil {
		message = "something went wrong when update"
		success = false
		return
	}

	data, _ = GetCart(userID)
	message = "product updated successfully"
	success = true
	return

}

func IsHasProductOnCart(userID primitive.ObjectID) (hasProduct bool, err error) {
	hasProduct = false

	cart, err := GetCart(userID)
	if err != nil {
		return
	}

	cartProductLength := len(*cart.Product)

	if cart.Product == nil {
		return
	}

	if cartProductLength <= 0 {
		return
	}

	hasProduct = true
	return
}

func ClearProductsInCart(userID primitive.ObjectID) (cart *models.Cart) {
	cartCollection := database.MI.DB.Collection("carts")
	var cartProduct []models.CartProduct = make([]models.CartProduct, 0)

	query := bson.D{{Key: "userid", Value: userID}}
	update := bson.M{"$set": bson.M{
		"product": cartProduct,
	}}

	cartCollection.FindOneAndUpdate(context.Background(), query, update).Decode(cart)
	return cart
}
