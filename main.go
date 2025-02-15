package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/harshdangi/go-rest-api-with-fiber/database"
	"github.com/harshdangi/go-rest-api-with-fiber/router"
)

func setupServer() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())

	app.Route("/api/", func(api fiber.Router) {
		router.SetupRoutes(api)
	})
	return app
}

func main() {
	if err := database.Connect(""); err != nil {
		log.Fatal(err)
	}
	app := setupServer()
	app.Listen(":3000")
}
