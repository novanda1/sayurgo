package config

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func WithSwagger(app *fiber.App) {
	app.Get("/swagger", swagger.Handler)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))
}
