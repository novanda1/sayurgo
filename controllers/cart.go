package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/models"
	"github.com/novanda1/sayurgo/services"
	"github.com/novanda1/sayurgo/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCart(c *fiber.Ctx) error {
	body := &models.Cart{}
	err := c.BodyParser(body)

	cart, err := services.CreateCart(*body)

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
	cart, err := services.GetCart(paramId)

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

	paramProductId := c.Params("productid")
	id, err := primitive.ObjectIDFromHex(paramProductId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": "false",
			"message": "wrong param id",
		})
	}

	cartProduct.ProductID = &paramProductId
	errors := cartProduct.Validate(*cartProduct)
	if errors != nil {
		return c.JSON(errors)
	}

	_, err = services.GetProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Product Not found",
			"error":   err,
		})
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

	cart, msg := services.AddProductToCart(c, *user.ID, cartProduct)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": msg,
		"data": fiber.Map{
			"cart": cart,
		},
	})
}
