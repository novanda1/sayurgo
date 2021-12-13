package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/services"
	"github.com/novanda1/sayurgo/models"
	"github.com/novanda1/sayurgo/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(c *fiber.Ctx) error {
	body := &models.Order{}
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	phone := utils.GetPhoneFromJWT(c)
	user, err := services.GetUserByPhone(phone)
	userID, _ := primitive.ObjectIDFromHex(*user.ID)
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
			"error":   err,
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"order": order,
		},
	})
}
