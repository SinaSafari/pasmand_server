package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"pasmand/internal/utils/jwt"
	"strings"
)

func AuthMiddleware(c *fiber.Ctx) error {
	h := c.Get("Authorization")

	if h == "" {
		return fiber.ErrUnauthorized
	}

	chunks := strings.Split(h, " ")
	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}

	phone, err := jwt.Verify(chunks[1])
	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Locals("phone", phone)
	return c.Next()
}
