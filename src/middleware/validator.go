package middleware

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

const (
	Body = iota
	Params
	Query
)

func Validate[T any](segment int) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		target := new(T)
		switch segment {
		case Body:
			ctx.BodyParser(target)
		case Params:
			ctx.ParamsParser(target)
		case Query:
			ctx.QueryParser(target)
		}
		formattedErrs := []string{}
		errs := validate.Struct(target)
		if errs != nil {
			for _, err := range errs.(validator.ValidationErrors) {
				formattedErrs = append(formattedErrs, fmt.Sprintf("%s failed on the '%s' tag against value '%s'", err.Field(), err.Tag(), err.Value()))
			}
		}
		if len(formattedErrs) > 0 {
			panic(fiber.NewError(fiber.StatusUnprocessableEntity, strings.Join(formattedErrs, " , ")))
		}
		return ctx.Next()
	}
}
