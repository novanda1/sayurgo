package services

import (
	"context"
	"fmt"
	"time"

	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveOtp(body models.Otp) (otp *models.Otp, err error) {
	otp, err = GetOtpByPhone(body.Phone)

	if err != nil {
		body.Exp = time.Now().Add(30 * time.Second)
		otp, err = CreateOtp(body)
		return
	}

	otp, err = ModifyOtpKey(body)

	return
}

func GetOtpByPhone(phone *string) (otp *models.Otp, err error) {
	otpCollection := config.MI.DB.Collection("otps")
	query := bson.M{"phone": phone}
	err = otpCollection.FindOne(context.Background(), query).Decode(&otp)
	return
}

func GetOtpByIDAfterInsert(otpID interface{}) (otp *models.Otp, err error) {
	otpCollection := config.MI.DB.Collection("otps")
	query := bson.M{"_id": otpID}

	otp = new(models.Otp)
	err = otpCollection.FindOne(context.Background(), query).Decode(otp)

	return
}

func CreateOtp(incomingOtp models.Otp) (otp *models.Otp, err error) {
	otpCollection := config.MI.DB.Collection("otps")
	result, err := otpCollection.InsertOne(context.TODO(), incomingOtp)

	if err != nil {
		return
	}

	otp, err = GetOtpByIDAfterInsert(result.InsertedID)
	return
}

func ModifyOtpKey(incomingOtp models.Otp) (otp *models.Otp, err error) {
	otpCollection := config.MI.DB.Collection("otps")
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
	otpCollection := config.MI.DB.Collection("otps")
	query := bson.M{"phone": phone}
	result := otpCollection.FindOneAndDelete(context.Background(), query)
	err = result.Err()

	fmt.Println(result)

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
	} else {
		verified = false
		DeleteOtp(&phone)
		return
	}
}
