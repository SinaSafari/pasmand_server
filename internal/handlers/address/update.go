package address

import "github.com/gofiber/fiber/v2"

func UpdateAddress(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
