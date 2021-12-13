package adminControllers

import (
	"github.com/gofiber/fiber/v2"
	adminServices "github.com/novanda1/sayurgo/admin/services"
	"github.com/novanda1/sayurgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChangeOrderStatus(c *fiber.Ctx) error {
	body := new(models.Order)
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	paramID := c.Params("id")
	orderID, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	order, err := adminServices.ChangeOrderStatus(orderID, models.OrderStatus(body.Status))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"order": order,
		},
	})
}
