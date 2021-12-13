package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/controllers"
)

func ProductRoute(route fiber.Router) {
	route.Get("", controllers.GetProducts)
	route.Get("/:id", controllers.GetProduct)
}
