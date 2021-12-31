package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
)

var ConfigDefault = cache.Config{
	Next:         nil,
	Expiration:   1 * time.Minute,
	CacheControl: false,
	KeyGenerator: func(c *fiber.Ctx) string {
		return utils.CopyString(c.Path())
	},
	ExpirationGenerator: nil,
	Storage:             nil,
}

func WithCache(app *fiber.App) {
	app.Use(cache.New(ConfigDefault))
}
