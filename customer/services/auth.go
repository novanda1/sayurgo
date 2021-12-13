package services

import (
	"time"

	"github.com/hgfischer/go-otp"
	"github.com/novanda1/sayurgo/models"

	"github.com/golang-jwt/jwt/v4"
)

var secret = `12345678901234567890`
var hotp = &otp.TOTP{Secret: secret, Length: 6, IsBase32Secret: true, Period: 60}

func Auth(body *models.User) (otp *models.Otp, err error) {
	otpkey := hotp.Get()
	otp, err = SaveOtp(body.Phone, &otpkey)
	return
}

func AuthVerif(phone *string, otpkey *string) bool {
	verified := VerifyOtp(phone, otpkey)
	return verified
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
