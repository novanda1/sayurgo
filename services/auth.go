package services

import (
	"time"

	"github.com/novanda1/sayurgo/models"

	"github.com/golang-jwt/jwt/v4"
)

func Auth(body *models.User) (token string, message string, user *models.User) {
	user, err := GetUserByPhone(*body.Phone)
	if err != nil {
		err, user = CreateUser(*body)

		if err != nil {
			return "", "failed createuser", user
		}

		token, _ := SignToken(user)

		return token, "successfully", user
	}

	token, _ = SignToken(user)

	return token, "successfully", user
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
