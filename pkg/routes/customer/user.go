package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/novanda1/sayurgo/app/controllers/customer"
)

func UserRoute(route fiber.Router) {
	route.Put("", controllers.UpdateMyProfile)
	route.Post("/address", controllers.AddMyAddress)
}
