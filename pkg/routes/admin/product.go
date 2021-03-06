package adminRoutes

import (
	"github.com/gofiber/fiber/v2"
	adminControllers "github.com/novanda1/sayurgo/app/controllers/admin"
)

func ProductRoute(route fiber.Router) {
	route.Post("", adminControllers.AdminCreateProduct)
	route.Put("/:id", adminControllers.AdminUpdateProduct)
	route.Delete("/:id", adminControllers.AdminDeleteProduct)
}
