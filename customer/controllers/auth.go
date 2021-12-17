package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/services"
	"github.com/novanda1/sayurgo/models"
)

type verifOtpParams struct {
	Otp   *string `json:"otp,omitempty" bson:"otp,omitempty"`
	Phone *string `json:"phone,omitempty" bson:"phone,omitempty"`
}

type requestOtpParams struct {
	Phone *string `json:"phone,omitempty" bson:"phone,omitempty"`
}

// Auth func request Authorization that return OTP code.
// @Description Request Authorization that return OTP code.
// @Summary Request OTP code
// @Tags Auth
// @Accept json
// @Produce json
// @Param phone body requestOtpParams true "Your Phone Number"
// @Success 200 {object} verifOtpParams
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
		return c.JSON(fiber.Map{"otp": nil, "success": false, "error": err.Error()})
	}

	return c.JSON(fiber.Map{"otp": otp.Otp, "success": true})
}

// Auth func verif Authorization that return JWT code.
// @Description verif Authorization that return JWT code.
// @Summary Verify OTP code and get JWT code
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body verifOtpParams true "Your Phone Number"
// @Success 200 {object} models.User
// @Router /api/verif [post]
func AuthVerif(c *fiber.Ctx) error {
	body := new(verifOtpParams)
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
