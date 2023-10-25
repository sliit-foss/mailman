package v1

import (
	"mailman/src/dto"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return c.JSON(dto.Response[*interface{}]{
		Message: "Login successful",
	})
}

