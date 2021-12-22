package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/novanda1/sayurgo/app/controllers/customer"
)

func AuthRoute(route fiber.Router) {
	route.Post("/request", controllers.Auth)
	route.Post("/verify", controllers.AuthVerif)
}
