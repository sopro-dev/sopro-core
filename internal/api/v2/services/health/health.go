package health

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	log.Println("[Health][Ping]")
	return c.SendString("Pong!")
}
