package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/novanda1/sayurgo/app/controllers/customer"
)

func CartRoute(route fiber.Router) {
	route.Get("", controllers.GetCart)
	route.Post("/", controllers.CreateCart)
	route.Post("/:productid", controllers.AddProductToCart)
	route.Delete("/:productid", controllers.DeleteProductFromCart)
	route.Put("/:productid", controllers.ChangeTotalProductInCart)
}
