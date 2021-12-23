package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	adminRoutes "github.com/novanda1/sayurgo/pkg/routes/admin"
	routes "github.com/novanda1/sayurgo/pkg/routes/customer"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	auth := app.Group("/auth")
	api := app.Group("/api").Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	routes.AuthRoute(auth)
	routes.ProductRoute(api.Group("/products"))
	routes.CartRoute(api.Group("/carts"))
	routes.OrderRoute(api.Group("/order"))
	routes.UserRoute(api.Group("/user"))
}

func SetupAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin").Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	adminRoutes.ProductRoute(admin.Group("/products"))
	adminRoutes.OrderRoute(admin.Group("/orders"))
}
