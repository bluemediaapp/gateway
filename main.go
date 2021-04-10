package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bluemediaapp/gateway/auth"
	"log"
	"os"
)

var (
	app    = fiber.New()
	config *Config
)

func main() {
	config = &Config{
		port:     os.Getenv("port"),
		mongoUri: os.Getenv("mongo_uri"),
	}

	app.Use(func(ctx *fiber.Ctx) error {
		userId, err := auth.RequireLogin(ctx)
		if err == nil {
			ctx.Locals("logged-on", true)
			ctx.Locals("user-id", userId)
		} else {
			ctx.Locals("logged-on", false)
		}
		return ctx.Next()
	})

	log.Fatal(app.Listen(config.port))
}