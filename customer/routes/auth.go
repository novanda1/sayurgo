package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/customer/controllers"
)

func AuthRoute(route fiber.Router) {
	route.Post("/login", controllers.Auth)
}
