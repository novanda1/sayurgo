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

// Auth func request Authorization that return OTP code.
// @Description Request Authorization that return OTP code.
// @Summary get OTP code
// @Tags Auth
// @Accept json
// @Produce json
// @Param phone body string true "Your Phone Number"
// @Success 200 {string} otp
// @Router /auth/login [post]
func Auth(c *fiber.Ctx) error {
	body := new(models.Otp)
	err := c.BodyParser(&body)

	if err != nil {
		c.JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
		})
	}

	errors := body.Validate(*body)
	if errors != nil {
		return c.JSON(errors)
	}

	otp, err := services.Auth(body)

	if err != nil {
		return c.JSON(fiber.Map{"otp": otp, "success": false, "error": err.Error()})
	}

	return c.JSON(fiber.Map{"otp": otp, "success": true})
}

// Auth func verif Authorization that return JWT code.
// @Description verif Authorization that return JWT code.
// @Summary get JWT code
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /api/verif [post]
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
