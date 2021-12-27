package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/app/models"
	"github.com/novanda1/sayurgo/app/services"
	"github.com/novanda1/sayurgo/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateMyProfile(c *fiber.Ctx) error {
	body := &models.User{}
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "wrong body",
			"error":   err.Error(),
		})
	}

	if body.Role != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "change role is not allowed",
		})
	}

	userid := utils.GetUseridFromJWT(c)
	primitiveUserid, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "try to relogin",
			"error":   err.Error(),
		})
	}

	updatedUser, err := services.UpdateUser(primitiveUserid, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "failed to update",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "updated successfully",
		"user":    updatedUser,
	})
}

func AddMyAddress(c *fiber.Ctx) error {
	body := &models.UserAddress{}
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "wrong body",
			"error":   err.Error(),
		})
	}

	userid := utils.GetUseridFromJWT(c)
	primitiveUserid, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "try to relogin",
			"error":   err.Error(),
		})
	}

	updatedUser, err := services.AddUserAddress(primitiveUserid, *body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "failed to update",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "updated successfully",
		"user":    updatedUser,
	})
}
