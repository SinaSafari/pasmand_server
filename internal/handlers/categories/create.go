package categories

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pasmand/internal/config/database"
	"pasmand/internal/models"
)

type CreateCategoryPayload struct {
	Title string `json:"title"`
}

func CreateCategory(c *fiber.Ctx) error {
	payload := new(CreateCategoryPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "body is not in correct format",
		})
	}

	category := models.Category{Title: payload.Title}
	result := database.DB.Create(&category)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": ""})
	}
	return c.JSON(fiber.Map{"data": category.ID})
}
