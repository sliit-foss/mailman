package system

import (
	"github.com/gofiber/fiber/v2"
	"mailman/src/dto"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(dto.Response[*interface{}]{
		Message: "Mailman up and running",
	})
}
