package services

import (
	"os"
	"time"

	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/pkg/utils"

	"github.com/golang-jwt/jwt/v4"
)

func Auth(body *models.Otp) (otp *models.Otp, err error) {
	otpkey := utils.GenerateOTP(6)
	body.Otp = &otpkey

	otp, err = SaveOtp(*body)
	if err != nil {
		return
	}

	otp, err = GetOtpByPhone(body.Phone)

	return
}

func AuthVerification(phone *string, otpkey *string) (result models.VerifOtpResult, fail error) {
	verified := VerifyOtp(*phone, *otpkey)
	result.Verified = &verified

	if verified {
		user, err := GetUserByPhone(*phone)
		if err != nil {
			userData := new(models.User)
			userData.Phone = phone
			user, err = CreateUser(*userData)
			token, _ := SignToken(user)

			result.User = user
			result.Token = &token

			err = nil

			return
		}

		token, err := SignToken(user)
		result.User = user
		result.Token = &token

		return
	}

	return
}

func SignToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"phone": user.Phone,
		"id":    user.ID,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return t, err
}
