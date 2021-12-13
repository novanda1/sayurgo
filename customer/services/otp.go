package services

import (
	"context"
	"time"

	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveOtp(phone *string, otpkey *string) (otp *models.Otp, err error) {
	_, err = GetOtpByPhone(phone)
	if err == nil {
		otp, err = ModifyOtpKey(phone, otpkey)
		return
	}

	otp = new(models.Otp)
	otp.Otp = otpkey
	otp.Phone = phone
	otp.Exp = time.Now().Local().Add(60 * time.Second)

	otp, err = CreateOtp(otp, phone)
	return
}

func GetOtpByPhone(phone *string) (otp *models.Otp, err error) {
	otpCollection := config.MI.DB.Collection("otps")
	query := bson.M{"phone": phone}
	err = otpCollection.FindOne(context.Background(), query).Decode(otp)
	return
}

func GetOtpByIDAfterInsert(otpID interface{}) (otp *models.Otp, err error) {
	otpCollection := config.MI.DB.Collection("otps")
	query := bson.M{"_id": otpID}
	err = otpCollection.FindOne(context.Background(), query).Decode(otp)
	return
}

func CreateOtp(otpParams *models.Otp, phone *string) (otp *models.Otp, err error) {
	otpCollection := config.MI.DB.Collection("otps")
	result, err := otpCollection.InsertOne(context.Background(), otp)
	if err != nil {
		return
	}

	otp, err = GetOtpByIDAfterInsert(result.InsertedID)
	return
}

func ModifyOtpKey(phone *string, newOtp *string) (otp *models.Otp, err error) {
	otpCollection := config.MI.DB.Collection("otps")
	otp, err = GetOtpByPhone(phone)
	if err != nil {
		return
	}

	query := bson.M{"phone": phone}
	update := bson.M{"$set": bson.M{"otp": newOtp}}
	err = otpCollection.FindOneAndUpdate(context.Background(), query, update).Decode(otp)

	return
}

func VerifyOtp(phone *string, otpkey *string) (verified bool) {
	verified = false
	otp, err := GetOtpByPhone(phone)
	if err != nil {
		return
	} else if otp.Otp == otpkey {
		verified = true
		return
	} else {
		return
	}
}
