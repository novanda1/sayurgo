package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/services"
	"github.com/novanda1/sayurgo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all product returned array of products.
// @Description Get all product returned array of products.
// @Summary Get All Products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /api/products [get]
func GetProducts(c *fiber.Ctx) error {
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

	products, hasNext, err := services.AllProducts(*options)
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
			"hasNext":  hasNext,
		},
	})
}

// Get product returned products object.
// @Description Get product returned products object.
// @Summary Get Product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Param id path string true "Product ID"
// @Router /api/products/{id} [get]
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
