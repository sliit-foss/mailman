package system

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	system := fiber.New()
	system.Get("/health", Health)
	system.Get("/memory", Memory)
	return system
}
