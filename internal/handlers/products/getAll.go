package products

import (
	"github.com/gofiber/fiber/v2"
	"pasmand/internal/config/database"
	"pasmand/internal/models"
)

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Model(&models.Product{}).Find(&products)
	return c.JSON(fiber.Map{"data": products})
}
