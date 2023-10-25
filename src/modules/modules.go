package modules

import (
	"github.com/gofiber/fiber/v2"
	"mailman/src/modules/auth"
	"mailman/src/modules/system"
)

func New() *fiber.App {
	modules := fiber.New()
	modules.Mount("/auth", auth.New())
	modules.Mount("/system", system.New())
	return modules
}
