package services

import (
	"time"

	"github.com/hgfischer/go-otp"
	"github.com/novanda1/sayurgo/models"

	"github.com/golang-jwt/jwt/v4"
)

var secret = `12345678901234567890`
var hotp = &otp.TOTP{Secret: secret, Length: 6, IsBase32Secret: true, Period: 60}

func Auth(body *models.User) (otpkey string, message string, user *models.User) {
	user, err := GetUserByPhone(*body.Phone)
	if err != nil {
		err, user = CreateUser(*body)

		if err != nil {
			return "", "failed createuser", user
		}
	}

	otpkey = hotp.Get()

	return otpkey, "successfully", user
}

func AuthVerif(otpkey *string) bool {
	verif := hotp.Verify(*otpkey)
	return verif
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
