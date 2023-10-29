package modules

import (
	"mailman/src/middleware"
	"mailman/src/modules/auth"
	"mailman/src/modules/system"
	"mailman/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	modules := fiber.New()

	modules.Mount("/auth", auth.New())

	modules.Mount("/system", system.New())

	modules.All("/*", middleware.Protect)

	modules.Mount("/users", users.New())

	return modules
}
