package adminRoutes

import (
	"github.com/gofiber/fiber/v2"
	adminControllers "github.com/novanda1/sayurgo/admin/controllers"
)

func AuthRoutes(route fiber.Router) {
	route.Post("", adminControllers.AdminAuth)
}
