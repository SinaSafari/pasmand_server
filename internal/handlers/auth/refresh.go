package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pasmand/internal/utils/jwt"
	"time"
)

type RefreshPayload struct {
	RefreshToken string `json:"refresh_token"`
}

func Refresh(c *fiber.Ctx) error {
	payload := new(RefreshPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": "body is not in correct format"})
	}
	phone, err := jwt.Verify(payload.RefreshToken)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "token is not valid",
		})
	}

	accessToken, err := jwt.GenerateJwtToken(fmt.Sprintf("%v", phone), time.Now().Add(5*time.Minute))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "err in jwt generation",
			"err":    err,
		})
	}

	return c.JSON(fiber.Map{"routeName": "Refresh", "access_token": accessToken})
}
