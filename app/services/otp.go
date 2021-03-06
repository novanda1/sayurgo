package services

import (
	"context"
	"time"

	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/platform/database"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveOtp(body models.Otp) (otp *models.Otp, err error) {
	_, err = GetOtpByPhone(body.Phone)
	if err != nil {
		body.Exp = time.Now().Add(2 * time.Minute)
		otp, err = CreateOtp(body)
		return
	}

	otp, err = ModifyOtpKey(body)

	return
}

func GetOtpByPhone(phone *string) (otp *models.Otp, err error) {
	otpCollection := database.MI.DB.Collection(models.OtpCollectionName)
	query := bson.M{"phone": phone}
	err = otpCollection.FindOne(context.Background(), query).Decode(&otp)
	return
}

func GetOtpByIDAfterInsert(otpID interface{}) (otp *models.Otp, err error) {
	otpCollection := database.MI.DB.Collection(models.OtpCollectionName)
	query := bson.M{"_id": otpID}

	otp = new(models.Otp)
	err = otpCollection.FindOne(context.Background(), query).Decode(otp)

	return
}

func CreateOtp(incomingOtp models.Otp) (otp *models.Otp, err error) {
	otpCollection := database.MI.DB.Collection(models.OtpCollectionName)
	result, err := otpCollection.InsertOne(context.TODO(), incomingOtp)

	if err != nil {
		return
	}

	otp, err = GetOtpByIDAfterInsert(result.InsertedID)
	return
}

func ModifyOtpKey(incomingOtp models.Otp) (otp *models.Otp, err error) {
	otpCollection := database.MI.DB.Collection(models.OtpCollectionName)
	otp, err = GetOtpByPhone(incomingOtp.Phone)
	if err != nil {
		return
	}

	query := bson.M{"phone": incomingOtp.Phone}
	update := bson.M{"$set": bson.M{"otp": incomingOtp.Otp}}
	err = otpCollection.FindOneAndUpdate(context.Background(), query, update).Decode(otp)

	return
}

func DeleteOtp(phone *string) (err error) {
	otpCollection := database.MI.DB.Collection(models.OtpCollectionName)
	query := bson.M{"phone": phone}
	result := otpCollection.FindOneAndDelete(context.Background(), query)
	err = result.Err()

	return
}

func VerifyOtp(phone string, otpkey string) (verified bool) {
	otp, err := GetOtpByPhone(&phone)
	if err != nil {
		return
	}

	now := time.Now()
	pass := now.After(otp.Exp)    // check wheter token is expired
	verified = otpkey == *otp.Otp // check wether token is same

	if verified && !pass {
		DeleteOtp(&phone)
		return
	} else if pass {
		DeleteOtp(&phone)
		return false
	} else {
		return false
	}
}
