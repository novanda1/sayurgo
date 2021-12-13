package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/controllers"
)

func OrderRoute(route fiber.Router) {
	route.Post("", controllers.CreateOrder)
}
