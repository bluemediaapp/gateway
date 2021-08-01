package main

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

var (
	client = &http.Client{}
)

func initCachedApi() {
	app.Get("/api/cached/user/:user_id", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("user_id")
		// Validate user id
		if !validateId(userId) {
			return ctx.Status(400).SendString("Invalid user-id.")
		}
		return forwardHttp(ctx, "http://data-outlet:5000/user/"+userId, nil)
	})
	app.Get("/api/cached/video-info/:video_id", func(ctx *fiber.Ctx) error {
		videoId := ctx.Params("video_id")
		// Validate video id
		if !validateId(videoId) {
			return ctx.Status(400).SendString("Invalid video-id.")
		}
		return forwardHttp(ctx, "http://data-outlet:5000/video/"+videoId, nil)
	})
	app.Get("/api/cached/video/:video_id", func(ctx *fiber.Ctx) error {
		videoId := ctx.Params("video_id")
		// Validate video id
		if !validateId(videoId) {
			return ctx.Status(400).SendString("Invalid video-id.")
		}
		return forwardHttp(ctx, "http://cdn:5000/videos/"+videoId, nil)
	})
	app.Get("/api/cached/avatar/:user_id", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("user_id")
		if !validateId(userId) {
			return ctx.Status(400).SendString("Invalid user-id.")
		}
		return forwardHttp(ctx, "http://avatargen:5000/"+userId, nil)
	})
}

func validateId(id string) bool {
	_, err := strconv.ParseInt(id, 10, 64)
	return err == nil
}

func forwardHttp(ctx *fiber.Ctx, url string, includeHeaders []string) error {
	request, err := http.NewRequest("GET", url, bytes.NewReader(ctx.Body()))
	if err != nil {
		return err
	}
	if includeHeaders != nil {
		for _, headerName := range includeHeaders {
			request.Header.Set(headerName, ctx.Get(headerName, ""))
		}
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	ctx.Set("Content-Type", response.Header.Get("Content-Type"))
	ctx.Status(response.StatusCode)
	return ctx.SendStream(response.Body)
}
