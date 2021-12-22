package adminRoutes

import (
	"github.com/gofiber/fiber/v2"
	adminControllers "github.com/novanda1/sayurgo/app/controllers/admin"
)

func AuthRoutes(route fiber.Router) {
	route.Post("", adminControllers.AdminAuth)
}
