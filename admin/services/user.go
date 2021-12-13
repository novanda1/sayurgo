package adminServices

import (
	"context"
	"time"

	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAdminByPhone(phone string) (admin *models.Admin, err error) {
	adminCollection := config.MI.DB.Collection("admin")
	admin = &models.Admin{}

	query := bson.D{{Key: "phone", Value: phone}}
	err = adminCollection.FindOne(nil, query).Decode(admin)

	return
}

func CreateAdmin(data models.Admin) (admin *models.Admin, err error) {
	adminCollection := config.MI.DB.Collection("admin")
	data.ID = nil
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := adminCollection.InsertOne(context.Background(), data)

	if err != nil {
		admin = &data
		return
	}

	admin = &models.Admin{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}
	err = adminCollection.FindOne(context.TODO(), query).Decode(admin)

	return
}
