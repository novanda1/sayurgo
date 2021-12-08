package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/controllers"
)

func CartRoute(route fiber.Router) {
	route.Get("/:id", controllers.GetCart)
	route.Post("/", controllers.CreateCart)
	route.Post("/:id", controllers.AddProductToCart)
}
