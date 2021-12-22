package config

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	adminRoutes "github.com/novanda1/sayurgo/pkg/routes/admin"
	routes "github.com/novanda1/sayurgo/pkg/routes/customer"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	api := app.Group("/api")

	routes.AuthRoute(api.Group("/auth"))
	routes.ProductRoute(api.Group("/products"))
	routes.CartRoute(api.Group("/carts").Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	})))
	routes.OrderRoute(api.Group("/order").Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	})))
}

func SetupAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin")

	adminRoutes.AuthRoutes(admin.Group("/auth"))
	adminRoutes.ProductRoute(admin.Group("/products").Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("adminsecret"),
	})))
	adminRoutes.OrderRoute(admin.Group("/orders").Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("adminsecret"),
	})))
}
