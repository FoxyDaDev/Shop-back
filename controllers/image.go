package controllers

import (
	"strconv"

	"gradientfit/backend/database"
	"gradientfit/backend/models"

	"github.com/gofiber/fiber/v2"
)

func AddImage(c *fiber.Ctx) error {
	pid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid product ID"})
	}

	var body struct {
		URL     string `json:"url"`
		AltText string `json:"altText"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	var prod models.Product
	if err := database.DB.First(&prod, pid).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "product not found"})
	}

	img := models.Image{
		ProductID: uint(pid),
		URL:       body.URL,
		AltText:   body.AltText,
	}
	if err := database.DB.Create(&img).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not add image"})
	}

	return c.Status(201).JSON(img)
}

func DeleteImage(c *fiber.Ctx) error {
	iid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid image ID"})
	}

	if err := database.DB.Delete(&models.Image{}, iid).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not delete image"})
	}

	return c.SendStatus(204)
}

func ListImagesForProduct(c *fiber.Ctx) error {
	pid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid product ID"})
	}

	var images []models.Image
	if err := database.DB.Where("product_id = ?", pid).Find(&images).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not fetch images"})
	}

	return c.JSON(images)
}
