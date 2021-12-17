package services

import (
	"time"

	"github.com/novanda1/sayurgo/models"
	"github.com/novanda1/sayurgo/utils"

	"github.com/golang-jwt/jwt/v4"
)

func Auth(body *models.Otp) (otp *models.Otp, err error) {
	otpkey := utils.GenerateOTP(6)
	body.Otp = &otpkey

	otp, err = SaveOtp(*body)

	return
}

func AuthVerification(phone *string, otpkey *string) (verified bool, user *models.User) {
	verified = VerifyOtp(phone, otpkey)
	userData := new(models.User)

	if verified {
		user, err := GetUserByPhone(*phone)
		if err != nil {
			userData := new(models.User)
			userData.Phone = phone
			user, err = CreateUser(*userData)
			return true, user
		}
		return true, user
	}

	return false, userData
}

func SignToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"phone": user.Phone,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	return t, err
}
