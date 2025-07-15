package controllers

import (
	"strconv"

	"gradientfit/backend/database"
	"gradientfit/backend/models"

	"github.com/gofiber/fiber/v2"
)

func AddVariant(c *fiber.Ctx) error {
	pid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid product ID"})
	}

	var body struct {
		Color string   `json:"color"`
		Size  string   `json:"size"`
		Price *float64 `json:"price"`
		Stock *int     `json:"stock"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	var prod models.Product
	if err := database.DB.First(&prod, pid).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "product not found"})
	}

	variant := models.Variant{
		ProductID: uint(pid),
		Color:     body.Color,
		Size:      body.Size,
	}
	if body.Price != nil {
		variant.Price = *body.Price
	}
	if body.Stock != nil {
		variant.Stock = *body.Stock
	}

	if err := database.DB.Create(&variant).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not create variant"})
	}
	return c.Status(201).JSON(variant)
}

func UpdateVariant(c *fiber.Ctx) error {
	vid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid variant ID"})
	}

	var variant models.Variant
	if err := database.DB.First(&variant, vid).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "variant not found"})
	}

	var body struct {
		Color *string  `json:"color"`
		Size  *string  `json:"size"`
		Price *float64 `json:"price"`
		Stock *int     `json:"stock"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if body.Color != nil {
		variant.Color = *body.Color
	}
	if body.Size != nil {
		variant.Size = *body.Size
	}
	if body.Price != nil {
		variant.Price = *body.Price
	}
	if body.Stock != nil {
		variant.Stock = *body.Stock
	}
	if err := database.DB.Save(&variant).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not update variant"})
	}
	return c.JSON(variant)
}

func DeleteVariant(c *fiber.Ctx) error {
	vid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid variant ID"})
	}

	if err := database.DB.Delete(&models.Variant{}, vid).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not delete variant"})
	}
	return c.SendStatus(204)
}

func ListVariantsByProduct(c *fiber.Ctx) error {
	pid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid product ID"})
	}

	var variants []models.Variant
	if err := database.DB.Where("product_id = ?", pid).Find(&variants).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not list variants"})
	}
	return c.JSON(variants)
}

func ListAllVariants(c *fiber.Ctx) error {
	var variants []models.Variant
	if err := database.DB.Find(&variants).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not list variants"})
	}
	return c.JSON(variants)
}
