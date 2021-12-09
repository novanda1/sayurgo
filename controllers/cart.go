package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/models"
	"github.com/novanda1/sayurgo/services"
	"github.com/novanda1/sayurgo/utils"
)

func CreateCart(c *fiber.Ctx) error {
	cart, err := services.CreateCart(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"product": cart,
		},
	})
}

func GetCart(c *fiber.Ctx) error {
	paramId := c.Params("id")
	cart, err := services.GetCart(c, paramId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "cart Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"cart": cart,
		},
	})
}

func AddProductToCart(c *fiber.Ctx) error {
	cartProduct := &models.CartProduct{}

	err := c.BodyParser(cartProduct)

	if err != nil {
		return err
	}

	phone := utils.GetPhoneFromJWT(c)

	user, err := services.GetUserByPhone(phone)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "user not found", // not authenticated
			"error":   err,
		})
	}

	cart, msg := services.AddProductToCart(user, cartProduct)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": msg,
		"data": fiber.Map{
			"cart": cart,
		},
	})
}
