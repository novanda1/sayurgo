package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/novanda1/sayurgo/app/controllers/customer"
)

func OrderRoute(route fiber.Router) {
	route.Post("", controllers.CreateOrder)
	route.Get("", controllers.GetOrders)
}
