package main

import (
	"fmt"
	"mailman/src/config"
	"mailman/src/global"
	"mailman/src/middleware"
	"mailman/src/modules"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {

	config.Load()

	app := fiber.New(fiber.Config{
		AppName:      "Mailman",
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(cors.New())

	app.Use(helmet.New())

	app.Use(recover.New())

	app.Use(requestid.New(requestid.Config{
		Header: global.CORRELATION_ID,
	}))

	app.Get("/metrics", monitor.New())

	app.Mount("/api", modules.New())

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.Port)))
}
