package main

import (
	"log"
	"os"

	"gradientfit/backend/database"
	"gradientfit/backend/models"
	"gradientfit/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	database.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Image{},
		&models.Variant{},
		&models.Cart{},
		&models.CartItem{},
	)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	routes.UserRoutes(app)
	routes.ProductRoutes(app)
	routes.ImageRoutes(app)
	routes.VariantRoutes(app)
	routes.CartRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(app.Listen(":" + port))
}
