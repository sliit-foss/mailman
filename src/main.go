package main

import (
	"fmt"
	"mailman/src/config"
	"mailman/src/global"
	"mailman/src/modules"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {

	config.Load()

	app := fiber.New(fiber.Config{
		AppName: "Mailman",
	})

	app.Use(cors.New())

	app.Use(helmet.New())

	app.Get("/metrics", monitor.New())

	app.Use(requestid.New(requestid.Config{
		Header: global.CORRELATION_ID,
	}))

	app.Mount("/api", modules.New())

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.Port)))
}
