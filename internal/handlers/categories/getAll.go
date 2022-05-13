package categories

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pasmand/internal/config/database"
	"pasmand/internal/models"
)

func GetAllCategories(c *fiber.Ctx) error {
	var data []models.Category
	results := database.DB.Model(&models.Category{}).Preload("Products").Find(&data)
	if results.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "something went wrong"})
	}
	return c.JSON(fiber.Map{"data": data})
}
