package v1

import (
	"github.com/gofiber/fiber/v2"
	"mailman/src/global"
)

func Create(c *fiber.Ctx) error {
	return c.JSON(global.Response[*interface{}]{
		Message: "User created successfully",
	})
}
