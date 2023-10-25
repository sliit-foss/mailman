package modules

import (
	"mailman/src/modules/auth"
	"mailman/src/modules/system"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
    modules := fiber.New()
	modules.Mount("/auth", auth.New())
    modules.Mount("/system", system.New())
    return modules
}