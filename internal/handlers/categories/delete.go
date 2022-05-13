package categories

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pasmand/internal/config/database"
	"pasmand/internal/models"
)

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{})
	}
	database.DB.Delete(&models.Category{}, id)
	return c.Status(http.StatusNoContent).JSON(fiber.Map{})
}
