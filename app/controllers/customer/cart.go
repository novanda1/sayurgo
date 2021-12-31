package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/app/services"
	"github.com/novanda1/sayurgo/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create an Empty cart for user
// @Description Create an Empty cart for user
// @Summary Create Empty Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {array} models.Cart
// @Router /api/carts [post]
func CreateCart(c *fiber.Ctx) error {
	body := &models.Cart{}
	userid := utils.GetUseridFromJWT(c)
	primitiveUserid, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	body.UserID = primitiveUserid
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

// Get cart data for specific user
// @Description Get cart data for specific user
// @Summary Get cart
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {array} models.Cart
// @Router /api/carts [get]
func GetCart(c *fiber.Ctx) error {
	useridContext := utils.GetUseridFromJWT(c)
	userID, _ := primitive.ObjectIDFromHex(useridContext)
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

// Add product to user cart
// @Description Add product to user cart
// @Summary Add product to cart
// @Tags Cart
// @Accept json
// @Produce json
// @Success 201 {array} models.Cart
// @Param productid path string true "Product ID"
// @Param body body models.CartProduct true "Product ID"
// @Router /api/carts/{productid} [post]
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
			"error":   err.Error(),
		})
	}

	userIdContext := utils.GetUseridFromJWT(c)
	userID, _ := primitive.ObjectIDFromHex(userIdContext)
	isExist := services.IsProductAlreadyExists(productID, userID)
	if isExist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "product already exist in your cart",
		})
	}

	cart, msg := services.AddProductToCart(userID, cartProduct)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": msg,
		"data": fiber.Map{
			"cart": cart,
		},
	})
}

// Delete specific product from user cart
// @Description Delete specific product from user cart
// @Summary Delete product from cart
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {array} models.Cart
// @Param productid path string true "Product ID"
// @Router /api/carts/{productid} [delete]
func DeleteProductFromCart(c *fiber.Ctx) error {
	paramId := c.Params("productid")
	productID, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	userIdContext := utils.GetUseridFromJWT(c)
	userID, err := primitive.ObjectIDFromHex(userIdContext)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	_, err = services.GetCart(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "cart Not found",
			"error":   err.Error(),
		})
	}

	message, success := services.DeleteProductFromCart(productID, userID)

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"success": success,
		"message": message,
	})
}

// Change Product Amount in specific product-cart
// @Description Change Product Amount in specific product-cart
// @Summary Change Product Amount in specific product-cart
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {array} models.Cart
// @Param productid path string true "Product ID"
// @Param amount body models.UpdateAmountCartProductParam true "Product Data"
// @Router /api/carts/{productid} [put]
func ChangeTotalProductInCart(c *fiber.Ctx) error {
	paramId := c.Params("productid")
	productID, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	userIdContext := utils.GetUseridFromJWT(c)
	userID, err := primitive.ObjectIDFromHex(userIdContext)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	_, err = services.GetCart(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "cart Not found",
		})
	}

	cartProduct := new(models.CartProduct)
	c.BodyParser(&cartProduct)

	message, success, data := services.ChangeTotalProductInCart(productID, userID, *cartProduct.TotalProduct)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": success,
		"message": message,
		"data":    data,
	})
}
