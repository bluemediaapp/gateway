package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func initCachedApi() {
	app.Get("/api/cached/user/:user_id", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("user_id")
		// Validate user id
		if !validateId(userId) {
			return ctx.Status(400).SendString("Invalid user-id.")
		}
		return forwardHttp(ctx, "http://data-outlet:5000/user/" + userId)
	})
	app.Get("/api/cached/video/:video_id", func(ctx *fiber.Ctx) error {
		videoId := ctx.Params("video_id")
		// Validate video id
		if !validateId(videoId) {
			return ctx.Status(400).SendString("Invalid video-id.")
		}
		return forwardHttp(ctx, "http://data-outlet:5000/video/" +videoId)
	})
}

func validateId(id string) bool {
	_, err := strconv.ParseInt(id, 10, 64)
	return err == nil
}

func forwardHttp(ctx *fiber.Ctx, url string) error {
	response, err := http.Get(url)
	if err != nil {
		_ = ctx.SendString(fmt.Sprint(err))
		return err
	}
	ctx.Set("Content-Type", response.Header.Get("Content-Type"))
	ctx.Status(response.StatusCode)
	return ctx.SendStream(response.Body)
}
