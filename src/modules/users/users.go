package users

import (
	"github.com/gofiber/fiber/v2"
	"mailman/src/middleware"
	"mailman/src/modules/users/api/v1"
)

func New() *fiber.App {
	users := fiber.New()
	users.All("/*", middleware.AdminProtect)
	users.Mount("/v1", v1.New())
	return users
}
