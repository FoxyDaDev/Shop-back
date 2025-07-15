package routes

import (
	"gradientfit/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	api := app.Group("/api/products")

	api.Get("/", controllers.ListProducts)
	api.Post("/", controllers.CreateProduct)
	api.Get("/:id", controllers.GetProductByID)
	api.Put("/:id", controllers.UpdateProduct)
	api.Delete("/:id", controllers.DeleteProduct)

	api.Get("/:id/variants", controllers.ListVariantsByProduct)
	api.Post("/:id/variants", controllers.AddVariant)
}
