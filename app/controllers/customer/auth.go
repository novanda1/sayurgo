package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/app/services"
	"github.com/novanda1/sayurgo/platform/whatsapp"
)

type verifOtpParams struct {
	Otp   *string `json:"otp,omitempty" bson:"otp,omitempty"`
	Phone *string `json:"phone,omitempty" bson:"phone,omitempty"`
}

type RequestOtpParams struct {
	Phone *string `json:"phone,omitempty" bson:"phone,omitempty"`
}

type VerifOtpResponse struct {
	Success *bool        `json:"success"`
	User    *models.User `json:"user"`
	Token   *string      `json:"token"`
}

// Auth func request Authorization that return OTP code.
// @Description Request Authorization that return OTP code.
// @Summary Request OTP code
// @Tags Auth
// @Accept json
// @Produce json
// @Param phone body requestOtpParams true "Your Phone Number"
// @Success 200 {object} verifOtpParams
// @Router /auth/request [post]
func Auth(c *fiber.Ctx) error {
	body := new(models.Otp)
	err := c.BodyParser(&body)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

	whatsapp.SendOtpCodeToWhatsapp(*otp.Phone, *otp.Otp)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// Auth func verif Authorization that return JWT token.
// @Description Verify OTP code.
// @Summary Verify OTP code and get JWT code
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body verifOtpParams true "Your Phone Number"
// @Success 200 {object} verifOtpResponse
// @Router /auth/verify [post]
func AuthVerif(c *fiber.Ctx) error {
	body := new(verifOtpParams)
	err := c.BodyParser(&body)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
		})
	}

	result, err := services.AuthVerification(body.Phone, body.Otp)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": result.Verified,
		"user":    result.User,
		"token":   result.Token,
	})
}
