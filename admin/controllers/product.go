package adminControllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/services"
	"github.com/novanda1/sayurgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AdminCreateProduct(c *fiber.Ctx) error {
	body := new(models.Product)
	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":  false,
			"messages": "failed to parse body",
		})
	}

	errors := body.Validate(*body)

	if errors != nil {
		return c.JSON(errors)
	}

	product, err := services.CreateProduct(*body)

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

func AdminUpdateProduct(c *fiber.Ctx) error {
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

	_, err = services.GetProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Product Not found",
			"error":   err,
		})
	}

	product, err := services.UpdateProduct(id, body)
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

func AdminDeleteProduct(c *fiber.Ctx) error {
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
