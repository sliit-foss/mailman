package users

import (
	"github.com/gofiber/fiber/v2"
	"mailman/src/modules/users/api/v1"
)

func New() *fiber.App {
	users := fiber.New()
	users.Mount("/v1", v1.New())
	return users
}
