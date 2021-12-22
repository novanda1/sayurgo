package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/novanda1/sayurgo/app/controllers/customer"
)

func ProductRoute(route fiber.Router) {
	route.Get("", controllers.GetProducts)
	route.Get("/:id", controllers.GetProduct)
}
