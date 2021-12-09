package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/models"

	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *fiber.Ctx) (token string, message string, user *models.User) {
	body := new(models.User)
	err := c.BodyParser(body)

	if body.Phone == nil {
		return "", "field phone is required", body
	}

	err, user = GetUserByPhone(c)
	if err != nil {
		err, user = CreateUser(c)

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
