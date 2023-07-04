package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pablodz/sopro/internal/api/v2/endpoints"
	"github.com/pablodz/sopro/internal/api/v2/services/health"
)

func HealthRoutes(app *fiber.App) {
	app.Get(endpoints.HEALTH, func(c *fiber.Ctx) error { return health.Ping(c) })
}
