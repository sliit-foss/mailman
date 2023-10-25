package modules

import (
	"mailman/src/modules/auth"
	"mailman/src/modules/system"
	"mailman/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	modules := fiber.New()
	modules.Mount("/auth", auth.New())
	modules.Mount("/users", users.New())
	modules.Mount("/system", system.New())
	return modules
}
