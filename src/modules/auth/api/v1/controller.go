package v1

import (
	"mailman/src/global"
	"mailman/src/modules/auth/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return c.JSON(global.Response[*interface{}]{
		Message: "Login successful",
	})
}

func Register(c *fiber.Ctx) error {
	payload := new(dto.RegisterReq)
	c.BodyParser(payload)
	res := registerUser(c, *payload)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.LoginRes]{
		Message: "Account created successfully",
		Data:    res,
	})
}
