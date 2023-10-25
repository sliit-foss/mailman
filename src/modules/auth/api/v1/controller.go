package v1

import (
	"github.com/gofiber/fiber/v2"
	"mailman/src/dto"
)

func Login(c *fiber.Ctx) error {
	return c.JSON(dto.Response[*interface{}]{
		Message: "Login successful",
	})
}
