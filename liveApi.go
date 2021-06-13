package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
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
		return forwardHttp(ctx, fmt.Sprintf("http://videos:5000/upload/%d", userId.(int64)), []string{"Content-Type", "Content-Length"})
	})

	app.Post("/api/live/delete-video/:video_id", func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("user-id")
		if userId == nil {
			return ctx.Status(403).SendString("Invalid login token")
		}

		videoId, err := strconv.ParseInt(ctx.Params("video_id"), 10, 64)

		if err != nil {
			return ctx.Status(400).SendString("Invalid video id")
		}

		return forwardHttp(ctx, fmt.Sprintf("http://videos:5000/delete/%d/%d", videoId, userId.(int64)), []string{"Content-Type", "Content-Length"})
	})

	app.Get("/api/live/set-watched/:video_id", func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("user-id")
		if userId == nil {
			return ctx.Status(403).SendString("You need to be logged in to use this")
		}

		videoId, err := strconv.ParseInt(ctx.Params("video_id"), 10, 64)
		if err != nil {
			return err
		}

		return forwardHttp(ctx, fmt.Sprintf("http://interactions:5000/watched/%d/%d", videoId, userId.(int64)), nil)
	})

	app.Get("/api/live/like/:video_id", func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("user-id")
		if userId == nil {
			return ctx.Status(403).SendString("You need to be logged in to use this")
		}

		videoId, err := strconv.ParseInt(ctx.Params("video_id"), 10, 64)
		if err != nil {
			return err
		}

		return forwardHttp(ctx, fmt.Sprintf("http://interactions:5000/like/%d/%d", videoId, userId.(int64)), nil)
	})

}
