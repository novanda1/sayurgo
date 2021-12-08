package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/services"
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
	cart, err := services.AddProductToCart(c)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "failed to add",
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
