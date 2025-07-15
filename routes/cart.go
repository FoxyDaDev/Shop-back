package routes

import (
	"gradientfit/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func CartRoutes(app *fiber.App) {
	api := app.Group("/api/cart")

	api.Get("/:userId", controllers.GetCart)
	api.Post("/:userId/add", controllers.AddToCart)
	api.Delete("/:userId/remove", controllers.RemoveFromCart)
	api.Post("/:userId/clear", controllers.ClearCart)
}
