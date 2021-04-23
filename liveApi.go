package main

import "github.com/gofiber/fiber/v2"

func initLiveApi() {
	app.Get("/api/live/recommended", func(ctx *fiber.Ctx) error {
		userId := ctx.Get("user-id", "")
		if userId == "" {
			return ctx.Status(403).SendString("Invalid login token")
		}
		return forwardHttp(ctx, "http://recommended:5000/user/"+userId, nil)
	})

	// Users
	app.Get("/api/live/login", func(ctx *fiber.Ctx) error {
		return forwardHttp(ctx, "http://users:5000/login", []string{"username", "password"})
	})
	app.Get("/api/live/register", func(ctx *fiber.Ctx) error {
		return forwardHttp(ctx, "http://users:5000/register", []string{"username", "password"})
	})
}
