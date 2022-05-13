package auth

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pasmand/internal/config/database"
	"pasmand/internal/models"
)

func Logout(c *fiber.Ctx) error {
	phone := c.Locals("phone")
	if phone == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "token is required",
		})
	}
	database.DB.Model(&models.User{}).Where("phone = ?", phone).Update("token", "")
	return c.JSON(fiber.Map{"routeName": "logout", "message": "ok"})
}
