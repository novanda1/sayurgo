package adminControllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/app/services"
	"github.com/novanda1/sayurgo/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create a brand new product.
// @Description Create a brand new product.
// @Summary Admin: Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Param body body models.Product false "Product"
// @Router /admin/products [post]
func AdminCreateProduct(c *fiber.Ctx) error {
	if !utils.AmIAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "401",
		})
	}

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

// Update some data in specific product.
// @Description Update some data in specific product.
// @Summary Admin: Update Product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Param body body models.Product false "Product"
// @Param id path int false "Product ID"
// @Router /admin/products/{id} [put]
func AdminUpdateProduct(c *fiber.Ctx) error {
	if !utils.AmIAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "401",
		})
	}

	paramID := c.Params("id")
	id, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "wrong param id",
		})
	}

	body := new(models.Product)
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
		})
	}

	_, err = services.GetProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Product Not found",
			"error":   err.Error(),
		})
	}

	product, err := services.UpdateProduct(id, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

// Delete specific product.
// @Description Delete specific product.
// @Summary Admin: Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int false "Product ID"
// @Router /admin/products/{id} [delete]
func AdminDeleteProduct(c *fiber.Ctx) error {
	if !utils.AmIAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "401",
		})
	}

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
