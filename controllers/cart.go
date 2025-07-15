package controllers

import (
	"strconv"

	"gradientfit/backend/database"
	"gradientfit/backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetCart(c *fiber.Ctx) error {
	uid, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	cart := models.Cart{UserID: uintPtr(uint(uid))}
	if err := database.DB.
		Preload("CartItems.Variant").
		FirstOrCreate(&cart, cart).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not load cart"})
	}

	return c.JSON(cart)
}

func AddToCart(c *fiber.Ctx) error {
	uid, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	var body struct {
		VariantID uint `json:"variantId"`
		Quantity  int  `json:"quantity"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	cart := models.Cart{UserID: uintPtr(uint(uid))}
	if err := database.DB.FirstOrCreate(&cart, cart).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not load cart"})
	}

	var item models.CartItem
	cond := models.CartItem{CartID: cart.ID, VariantID: body.VariantID}
	if err := database.DB.First(&item, cond).Error; err == nil {
		item.Quantity += body.Quantity
		database.DB.Save(&item)
	} else {
		item = models.CartItem{
			CartID:    cart.ID,
			VariantID: body.VariantID,
			Quantity:  body.Quantity,
		}
		database.DB.Create(&item)
	}

	return GetCart(c)
}

func RemoveFromCart(c *fiber.Ctx) error {
	uid, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	var body struct {
		VariantID uint `json:"variantId"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", uid).First(&cart).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "cart not found"})
	}

	if err := database.DB.
		Where("cart_id = ? AND variant_id = ?", cart.ID, body.VariantID).
		Delete(&models.CartItem{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not remove item"})
	}

	return GetCart(c)
}

func ClearCart(c *fiber.Ctx) error {
	uid, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", uid).First(&cart).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "cart not found"})
	}

	if err := database.DB.
		Where("cart_id = ?", cart.ID).
		Delete(&models.CartItem{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not clear cart"})
	}

	return c.JSON(fiber.Map{"message": "cart cleared"})
}

func uintPtr(u uint) *uint {
	return &u
}
