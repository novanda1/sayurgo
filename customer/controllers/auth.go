package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/services"
	"github.com/novanda1/sayurgo/models"
)

type verifParams struct {
	Otp *string `json:"otp,omitempty" bson:"otp,omitempty"`
}

func Auth(c *fiber.Ctx) error {
	body := new(models.User)
	err := c.BodyParser(&body)

	if err != nil {
		c.JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
		})
	}

	errors := body.AuthDtoValidate(*body)
	if errors != nil {
		return c.JSON(errors)
	}

	token, message, user := services.Auth(body)
	return c.JSON(fiber.Map{"token": token, "message": message, "user": user})
}

func AuthVerif(c *fiber.Ctx) error {
	body := new(verifParams)
	err := c.BodyParser(&body)

	if err != nil {
		c.JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
		})
	}

	verif := services.AuthVerif(body.Otp)
	return c.JSON(fiber.Map{"verif": verif})
}
