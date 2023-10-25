package main

import (
	"mailman/src/modules"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
    app := fiber.New(fiber.Config{
		AppName: "Mailman",
		EnablePrintRoutes: true,
	})
	app.Mount("/api", modules.New())
    log.Fatal(app.Listen(":3001"))
}