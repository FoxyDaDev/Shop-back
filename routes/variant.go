package routes

import (
	"gradientfit/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func VariantRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/products/:id/variants", controllers.AddVariant)

	api.Put("/variants/:id", controllers.UpdateVariant)

	api.Delete("/variants/:id", controllers.DeleteVariant)
	api.Get("/products/:id/variants", controllers.ListVariantsByProduct)
	api.Get("/variants", controllers.ListAllVariants)
}
