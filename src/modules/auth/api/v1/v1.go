package v1

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	v1 := fiber.New()
	v1.Post("/login", Login)
	return v1
}
