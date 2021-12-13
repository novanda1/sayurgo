package adminControllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	adminServices "github.com/novanda1/sayurgo/admin/services"
	"github.com/novanda1/sayurgo/models"
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

	token, message, user := adminServices.Auth(body)
	return c.JSON(fiber.Map{"token": token, "message": message, "user": user})
}
