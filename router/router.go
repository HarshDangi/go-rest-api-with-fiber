package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harshdangi/go-rest-api-with-fiber/handler"
)

func SetupRoutes(app fiber.Router) {
	app.Get("/", handler.GetAllProducts)
	app.Get("/:id", handler.GetSingleProduct)
	app.Post("/", handler.CreateProduct)
	app.Delete("/:id", handler.DeleteProduct)
}
