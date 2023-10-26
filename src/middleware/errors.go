package middleware

import (
	"errors"
	"mailman/src/global"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Just patching things up. This'll be over in a jiffy."
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}
	log.Error("Request error:", err.Error())
	return ctx.Status(code).JSON(global.Response[*interface{}]{
		Message: message,
	})
}
