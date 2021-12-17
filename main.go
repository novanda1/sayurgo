package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
	adminRoutes "github.com/novanda1/sayurgo/admin/routes"
	"github.com/novanda1/sayurgo/config"
	"github.com/novanda1/sayurgo/customer/routes"
	_ "github.com/novanda1/sayurgo/docs"

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

// @title SayurGO API
// @version 0.0.1
// @description The Sayurmax REST API built with GO

// @contact.name API Support
// @contact.url https://github.com/novanda1
// @contact.email novandaahsan1@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	app := fiber.New()
	app.Use(helmet.New())

	app.Get("/swagger", swagger.Handler)              // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:          "doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

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
