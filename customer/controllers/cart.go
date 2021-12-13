package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/services"
	"github.com/novanda1/sayurgo/models"
	"github.com/novanda1/sayurgo/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCart(c *fiber.Ctx) error {
	body := &models.Cart{}
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

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
	phone := utils.GetPhoneFromJWT(c)
	user, err := services.GetUserByPhone(phone)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	userID, _ := primitive.ObjectIDFromHex(*user.ID)
	cart, err := services.GetCart(userID)

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
	productID, err := primitive.ObjectIDFromHex(paramProductId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": "false",
			"message": "wrong param id",
		})
	}

	cartProduct.ProductID = productID
	errors := cartProduct.Validate(*cartProduct)
	if errors != nil {
		return c.JSON(errors)
	}

	_, err = services.GetProduct(productID)
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

	userID, _ := primitive.ObjectIDFromHex(*user.ID)
	isExist := services.IsProductAlreadyExists(productID, userID)
	if isExist {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "product already exist in your cart",
		})
	}

	cart, msg := services.AddProductToCart(userID, cartProduct)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": msg,
		"data": fiber.Map{
			"cart": cart,
		},
	})
}

func DeleteProductFromCart(c *fiber.Ctx) error {
	paramId := c.Params("productid")
	productID, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
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

	userID, err := primitive.ObjectIDFromHex(*user.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	_, err = services.GetCart(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "cart Not found",
			"error":   err,
		})
	}

	message, success := services.DeleteProductFromCart(productID, userID)

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": success,
		"message": message,
		"error":   err,
	})
}

func ChangeTotalProductInCart(c *fiber.Ctx) error {
	paramId := c.Params("productid")
	productID, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
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
	userID, err := primitive.ObjectIDFromHex(*user.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	_, err = services.GetCart(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "cart Not found",
			"error":   err,
		})
	}

	cartProduct := new(models.CartProduct)
	c.BodyParser(&cartProduct)

	message, success, data := services.ChangeTotalProductInCart(productID, userID, *cartProduct.TotalProduct)

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": success,
		"message": message,
		"data":    data,
	})
}
