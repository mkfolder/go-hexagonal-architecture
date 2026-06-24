package main

import (
	"github.com/gofiber/fiber/v3"
	"mkfolder.dev/wire-playground/internal/bootstrap"
)

func main() {
	app := fiber.New()
	router := app.Group("/api")

	router.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	bootstrap.InitializeContainer(router)
	app.Listen(":5000")
}
