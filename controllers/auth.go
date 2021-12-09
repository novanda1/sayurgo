package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/services"
)

func Auth(c *fiber.Ctx) error {
	token, message, user := services.Auth(c)

	return c.JSON(fiber.Map{"token": token, "message": message, "user": user})
}
