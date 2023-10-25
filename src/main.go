package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"mailman/src/modules"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:           "Mailman",
		EnablePrintRoutes: true,
	})
	app.Mount("/api", modules.New())
	log.Fatal(app.Listen(":3001"))
}
