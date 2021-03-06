package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/app/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all product returned array of products.
// @Description Get all product returned array of products.
// @Summary Get All Products
// @Tags Product
// @Accept json
// @Produce json
// @Param limit query int64 true "Limit"
// @Param lastid query string false "LastId"
// @Success 200 {array} models.GetAllProductsResponse
// @Router /api/products [get]
func GetProducts(c *fiber.Ctx) error {
	options := new(models.GetAllProductsParams)

	page := c.Query("lastid")
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	options.Limit = limit
	options.LastId = &page
	errors := options.Validate(*options)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
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
