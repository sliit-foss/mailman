package main

import (
	"fmt"
	"mailman/src/config"
	"mailman/src/database"
	"mailman/src/global"
	"mailman/src/middleware"
	"mailman/src/modules"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {

	config.Load()

	database.Connect()

	app := fiber.New(fiber.Config{
		AppName:           "Mailman",
		EnablePrintRoutes: true,
		ErrorHandler:      middleware.ErrorHandler,
	})

	app.Hooks().OnShutdown(database.Disconnect)

	app.Use(cors.New())

	app.Use(helmet.New())

	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(global.Response[*interface{}]{
				Message: "Too many requests, please try again later",
			})
		},
	}))

	app.Use(recover.New())

	app.Get("/metrics", monitor.New())

	app.Use(requestid.New(requestid.Config{
		Header: global.HeaderXCorrelationID,
	}))

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	app.Mount("/api", modules.New())

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.Port)))
}
