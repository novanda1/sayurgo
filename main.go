package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/routes"

	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	api := app.Group("/api")

	routes.AuthRoute(api.Group("/auth"))
	routes.ProductRoute(api.Group("/products"))
	routes.CartRoute(api.Group("/carts"))

}

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	setupRoutes(app)

	err = app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}
