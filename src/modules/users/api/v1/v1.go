package v1

import (
	"github.com/gofiber/fiber/v2"
	m "mailman/src/middleware"
	"mailman/src/modules/users/api/v1/dto"
)

func New() *fiber.App {
	v1 := fiber.New()
	v1.Post("/", m.Validate[dto.CreateUserReq](m.Body), Create)
	return v1
}
