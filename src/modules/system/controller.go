package system

import (
	"mailman/src/dto"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(dto.Response[*interface{}]{
		Message: "Mailman up and running",
	})
}

