package categories

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pasmand/internal/config/database"
	"pasmand/internal/models"
)

type UpdateCategoryPayload struct {
	Title string `json:"title"`
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{})
	}
	payload := new(UpdateCategoryPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "body is not in correct format",
		})
	}
	database.DB.Model(&models.Category{}).Where("id = ?", id).Updates(models.Category{Title: payload.Title})
	return c.Status(http.StatusAccepted).JSON(fiber.Map{})
}
