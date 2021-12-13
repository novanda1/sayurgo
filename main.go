package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	adminRoutes "github.com/novanda1/sayurgo/admin/routes"
	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/customer/routes"

	"github.com/joho/godotenv"

	jwtware "github.com/gofiber/jwt/v3"
)

func setupRoutes(app *fiber.App) {
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

func setupAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin")

	adminRoutes.AuthRoutes(admin.Group("/auth"))
	adminRoutes.ProductRoute(admin.Group("/products").Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("adminsecret"),
	})))
	adminRoutes.OrderRoute(admin.Group("/orders").Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("adminsecret"),
	})))
}

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	setupAdminRoutes(app)
	setupRoutes(app)

	err = app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}
