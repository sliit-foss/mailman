package system

import (
	"github.com/gofiber/fiber/v2"
	"mailman/src/global"
	"mailman/src/modules/system/dto"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(global.Response[*interface{}]{
		Message: "Mailman up and running",
	})
}

func Memory(c *fiber.Ctx) error {
	return c.JSON(global.Response[dto.MemStats]{
		Message: "Memory usage retrieved",
		Data:    GetMemoryUsage(),
	})
}
