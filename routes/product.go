package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/controllers"
)

func ProductRoute(route fiber.Router) {
	route.Get("", controllers.GetProducts)
	route.Post("", controllers.CreateProduct)
	route.Put("/:id", controllers.UpdateProduct)
	route.Delete("/:id", controllers.DeleteProduct)
	route.Get("/:id", controllers.GetProduct)
}
