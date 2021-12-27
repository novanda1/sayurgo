package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/app/services"
	"github.com/novanda1/sayurgo/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create an Order from Cart and delete all cart-product
// @Description Create an Order from Cart and delete all cart-product
// @Summary Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Success 201 {object} models.Order
// @Param order body models.Order true "Order"
// @Router /api/order/ [post]
func CreateOrder(c *fiber.Ctx) error {
	body := &models.Order{}
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	userIdContext := utils.GetUseridFromJWT(c)
	userID, _ := primitive.ObjectIDFromHex(userIdContext)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	hasProducts, err := services.IsHasProductOnCart(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if !hasProducts {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "you didnt have any products in your cart",
		})
	}

	order, err := services.CreateOrder(body, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"order": order,
		},
	})
}

// Get order data from current userid
// @Description Get order data from current userid
// @Summary Get Order
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {array} models.Order
// @Router /api/order/ [get]
func GetMyOrders(c *fiber.Ctx) error {
	useridString := utils.GetUseridFromJWT(c)
	userID, err := primitive.ObjectIDFromHex(useridString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
		})
	}

	order, err := services.GetOrdersByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    order,
	})
}
