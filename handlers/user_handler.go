package handlers

import (
	"fiber-generic/models"
	"fiber-generic/repositories"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Repo *repositories.GenericRepository[models.User]
}

// GetAllUsers handler untuk mendapatkan semua user
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.Repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// GetUserByID handler untuk mendapatkan user berdasarkan ID
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var uid uint
	fmt.Sscanf(id, "%d", &uid) // Konversi string ke uint

	user, err := h.Repo.GetByID(uid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// CreateUser handler untuk membuat user baru
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.Repo.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
