package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pablodz/sopro/internal/api/v2/endpoints"
	"github.com/pablodz/sopro/internal/api/v2/services/transcoding"
)

func TranscodingRoutes(app *fiber.App) {
	app.Post(endpoints.TRANSCODING, func(c *fiber.Ctx) error { return transcoding.TranscodeFile(c) })
}
