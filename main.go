package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
	_ "github.com/novanda1/sayurgo/docs"
	config "github.com/novanda1/sayurgo/pkg/configs"
	"github.com/novanda1/sayurgo/platform/database"

	"github.com/joho/godotenv"
)

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

	database.ConnectDB()

	config.SetupAdminRoutes(app)
	config.SetupRoutes(app)

	err = app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}
