package v1

import (
	"github.com/gofiber/fiber/v2"
	m "mailman/src/middleware"
	"mailman/src/modules/auth/api/v1/dto"
)

func New() *fiber.App {
	v1 := fiber.New()
	v1.Post("/login", m.Validate[dto.LoginReq](m.Body), Login)
	return v1
}
