package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
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
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	routes.ProductRoute(api.Group("/products"))
	api.Get("", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.SendString("Welcome " + name)
	})
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
