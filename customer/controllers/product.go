package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProducts(c *fiber.Ctx) error {
	products, err := services.AllProducts()

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
			"products": products,
		},
	})
}

func GetProduct(c *fiber.Ctx) error {
	paramId := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": "false",
			"message": "wrong param id",
		})
	}

	product, err := services.GetProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Product Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"product": product,
		},
	})
}
