package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/models"
	"github.com/novanda1/sayurgo/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProducts(c *fiber.Ctx) error {
	err, products := services.AllProducts()

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

func CreateProduct(c *fiber.Ctx) error {
	body := new(models.Product)
	c.BodyParser(&body)

	err, product := services.CreateProduct(*body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"product": product,
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

	err, product := services.GetProduct(id)

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

func UpdateProduct(c *fiber.Ctx) error {
	paramID := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramID)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "wrong param id",
		})
	}

	body := new(models.Product)
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
		})
	}

	err, product := services.UpdateProduct(id, body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
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

func DeleteProduct(c *fiber.Ctx) error {
	err := services.DeleteProduct(c)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Product Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete product",
			"error":   err,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
