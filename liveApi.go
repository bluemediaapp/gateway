package main

import "github.com/gofiber/fiber/v2"

func initLiveApi() {
	app.Get("/api/live/recommended", func(ctx *fiber.Ctx) error {
		userId := ctx.Get("user-id", "")
		if userId == "" {
			return ctx.Status(403).SendString("Invalid login token")
		}
		return forwardHttp(ctx, "http://recommended:5000/user/"+userId)
	})
}
