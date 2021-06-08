package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func initLiveApi() {
	app.Get("/api/live/recommended", func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("user-id")
		if userId == nil {
			return ctx.Status(403).SendString("Invalid login token")
		}
		return forwardHttp(ctx, fmt.Sprintf("http://recommended:5000/user/%d", userId.(int64)), nil)
	})

	// Users
	app.Get("/api/live/login", func(ctx *fiber.Ctx) error {
		return forwardHttp(ctx, "http://users:5000/login", []string{"username", "password"})
	})
	app.Get("/api/live/register", func(ctx *fiber.Ctx) error {
		return forwardHttp(ctx, "http://users:5000/register", []string{"username", "password"})
	})
	app.Get("/api/live/upload", func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("user-id")
		if userId == nil {
			return ctx.Status(403).SendString("Invalid login token")
		}
		return forwardHttp(ctx, fmt.Sprintf("http://interactions:5000/upload/%d", userId.(int64)), []string{"Content-Type", "Content-Length"})
	})
}
