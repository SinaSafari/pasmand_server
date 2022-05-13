package categories

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pasmand/internal/config/database"
	"pasmand/internal/models"
)

func GetOneCategory(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{})
	}
	var category models.Category
	database.DB.Find(&category, id)

	if category.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "category not found"})
	}
	return c.JSON(fiber.Map{"data": category})
}
