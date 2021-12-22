package adminControllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/app/services"
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

// Get all of order
// @Description Get all of order
// @Summary Get Order
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {array} models.Order
// @Router /api/order/ [get]
func GetOrders(c *fiber.Ctx) error {
	options := new(models.GetAllProductsParams)

	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	options.Limit = limit
	options.Page = page
	errors := options.Validate(*options)
	if errors != nil {
		return c.JSON(errors)
	}

	orders, hasNext, err := services.GetAllOrders(models.GetAllOrdersParams(*options))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"orders":  orders,
			"hasNext": hasNext,
		},
	})
}
