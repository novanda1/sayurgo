package adminControllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/app/services"
)

func AdminAuth(c *fiber.Ctx) error {
	body := new(models.Admin)
	err := c.BodyParser(&body)

	fmt.Println()

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
		})
	}

	errors := body.Validate(*body)
	if errors != nil {
		return c.JSON(errors)
	}

	token, message, user := services.AdminAuth(body)
	return c.JSON(fiber.Map{"token": token, "message": message, "user": user})
}
