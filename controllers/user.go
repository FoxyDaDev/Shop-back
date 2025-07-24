package controllers

import (
	"strconv"

	"gradientfit/backend/database"
	"gradientfit/backend/models"
	"gradientfit/backend/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}

func SignUp(c *fiber.Ctx) error {
	type request struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not hash password"})
	}

	user := models.User{
		Name:     body.Name,
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not create user"})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
	})
}

func Login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	if err := CheckPassword(user.Password, body.Password); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not generate token"})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":       user.ID,
			"name":     user.Name,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	var user models.User
	if err := database.DB.
		Select("id", "created_at", "updated_at", "name", "username", "email").
		First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(user)
}
