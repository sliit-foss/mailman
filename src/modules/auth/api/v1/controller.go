package v1

import (
	"github.com/gofiber/fiber/v2"
	"mailman/src/global"
)

func Login(c *fiber.Ctx) error {
	return c.JSON(global.Response[*interface{}]{
		Message: "Login successful",
	})
}
