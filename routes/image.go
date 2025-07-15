package routes

import (
	"gradientfit/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ImageRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/products/:id/images", controllers.AddImage)
	api.Delete("/images/:id", controllers.DeleteImage)
	api.Get("/products/:id/images", controllers.ListImagesForProduct)
}
