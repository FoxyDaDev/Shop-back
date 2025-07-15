package routes

import (
	"gradientfit/backend/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var jwtSecret = []byte("supersecretkey")

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/signup", controllers.SignUp)
	api.Post("/login", controllers.Login)

	protected := api.Group("/users")

	protected.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtSecret,
	}))

	protected.Get("/:id", controllers.GetUser)
}
