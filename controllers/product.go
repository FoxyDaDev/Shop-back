package controllers

import (
	"strconv"

	"gradientfit/backend/database"
	"gradientfit/backend/models"

	"github.com/gofiber/fiber/v2"
)

func ListProducts(c *fiber.Ctx) error {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not list products"})
	}
	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var payload struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		MainImage   string  `json:"mainImage"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	prod := models.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		MainImage:   payload.MainImage,
	}
	if err := database.DB.Create(&prod).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not create product"})
	}
	return c.Status(201).JSON(prod)
}

func GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid product ID"})
	}

	var prod models.Product
	if err := database.DB.First(&prod, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "product not found"})
	}
	return c.JSON(prod)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid product ID"})
	}

	var prod models.Product
	if err := database.DB.First(&prod, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "product not found"})
	}

	var payload struct {
		Name        *string  `json:"name"`
		Description *string  `json:"description"`
		Price       *float64 `json:"price"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	if payload.Name != nil {
		prod.Name = *payload.Name
	}
	if payload.Description != nil {
		prod.Description = *payload.Description
	}
	if payload.Price != nil {
		prod.Price = *payload.Price
	}

	if err := database.DB.Save(&prod).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not update product"})
	}
	return c.JSON(prod)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid product ID"})
	}
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not delete product"})
	}
	return c.SendStatus(204)
}
