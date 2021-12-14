package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/services"
	"github.com/novanda1/sayurgo/models"
)

type verifParams struct {
	Otp   *string `json:"otp,omitempty" bson:"otp,omitempty"`
	Phone *string `json:"phone,omitempty" bson:"phone,omitempty"`
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

	otp, err := services.Auth(body)
	return c.JSON(fiber.Map{"otp": otp, "error": err.Error()})
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

	verified, user := services.AuthVerification(body.Phone, body.Otp)
	return c.JSON(fiber.Map{
		"success": verified,
		"user":    user,
	})
}
