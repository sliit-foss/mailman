package v1;

import (
    "github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
    v1 := fiber.New()
    v1.Get("/login", Login)
    return v1
}