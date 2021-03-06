package adminRoutes

import (
	"github.com/gofiber/fiber/v2"
	adminControllers "github.com/novanda1/sayurgo/app/controllers/admin"
)

func OrderRoute(route fiber.Router) {
	route.Put("/:id", adminControllers.ChangeOrderStatus)
	route.Get("", adminControllers.GetOrders)
}
