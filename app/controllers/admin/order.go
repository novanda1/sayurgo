package adminControllers

import (
	"strconv"

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
	if !utils.AmIAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "401",
		})
	}

	body := new(models.Order)
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	paramID := c.Params("id")
	orderID, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	order, err := services.ChangeOrderStatus(orderID, models.OrderStatus(body.Status))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
// @Description Get all order of users
// @Summary Admin: Get Orders of all user
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {array} models.OrderResponse
// @Router /admin/order/ [get]
func GetOrders(c *fiber.Ctx) error {
	if !utils.AmIAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "401",
		})
	}

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
