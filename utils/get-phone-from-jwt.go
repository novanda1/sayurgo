package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetPhoneFromJWT(c *fiber.Ctx) (phone string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	phone = claims["phone"].(string)

	return phone
}
