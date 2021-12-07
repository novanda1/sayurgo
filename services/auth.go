package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/models"

	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *fiber.Ctx) (string, error) {
	// err, user := GetUserByPhone(c)
	// if err != nil {
	// 	fmt.Println(err)
	// 	token, err := SignToken(user)

	// 	return token, err
	// }

	err, user := CreateUser(c)
	token, err := SignToken(user)

	return token, err
}

func SignToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"name": user.DisplayName,
		"id":   user.ID,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	return t, err
}
