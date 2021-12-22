package adminControllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/app/services"
	"github.com/novanda1/sayurgo/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Admin-Only: Change OrderStatus on User-Order
// @Description Admin-Only: Change OrderStatus on User-Order
// @Summary Admin-Only: Change Order Status
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {array} models.Order
// @Param order body models.Order true "Order"
// @Router /admin/order/ [put]
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

	order, err := services.ChangeOrderStatus(orderID, models.OrderStatus(body.Status))
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

// Get order data from userid
// @Description Get order data from userid
// @Summary Get Order
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {array} models.Order
// @Router /api/order/ [get]
func GetOrders(c *fiber.Ctx) error {
	useridString := utils.GetUseridFromJWT(c)
	userID, err := primitive.ObjectIDFromHex(useridString)
	if err != nil {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
		})
	}

	order, err := services.GetOrdersByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    order,
	})
}
