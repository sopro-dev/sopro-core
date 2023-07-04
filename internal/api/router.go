package api

import (
	"log"
	"os"

	sentryfiber "github.com/aldy505/sentry-fiber"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pablodz/sopro/internal/api/v2/routes"
)

func HandleRequest() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		log.Fatalln("sentry initialization failed")
	}

	log.Printf("Fiber cold start")
	app := fiber.New()

	// Middleware: cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "authorization,Content-Type",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))

	// Middleware: sentry
	app.Use(sentryfiber.New(sentryfiber.Options{
		WaitForDelivery: false, // don't block lambda execution
	}))

	// Middleare: Logger
	app.Use(logger.New(logger.Config{
		Done: func(c *fiber.Ctx, logString []byte) {
			if c.Response().StatusCode() != fiber.StatusOK {
				sentry.CaptureMessage(string(logString))
			}
		},
	}))

	// Middleware: Recover
	app.Use(recover.New())

	app = Router(app)

	log.Fatal(app.Listen(":9898"))
}

func Router(app *fiber.App) *fiber.App {
	routes.HealthRoutes(app)
	routes.TranscodingRoutes(app)
	return app
}
