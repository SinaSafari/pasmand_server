package auth

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pasmand/internal/config/database"
	"pasmand/internal/config/redis"
	"pasmand/internal/models"
	"pasmand/internal/utils/jwt"
	"time"
)

type LoginPayload struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func Login(c *fiber.Ctx) error {
	ctx := context.Background()
	payload := new(LoginPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": "body is not in correct format"})
	}

	opt, err := redis.RedisClient.Get(ctx, payload.Phone).Result()
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(fiber.Map{
			"status": "redis server not available",
			"data":   err,
		})
	}

	if opt != payload.Code {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status": "code is not correct",
		})
	}

	accessTokenExpirationTime := time.Now().Add(5 * time.Minute)
	refreshTokenExpirationTime := time.Now().Add(time.Minute * 131400) // 3 months

	accessToken, err := jwt.GenerateJwtToken(payload.Phone, accessTokenExpirationTime)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "err in jwt generation",
			"err":    err,
		})
	}

	refreshToken, err := jwt.GenerateJwtToken(payload.Phone, refreshTokenExpirationTime)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "err in jwt generation",
			"err":    err,
		})
	}

	database.DB.Model(&models.User{}).Where("phone = ?", payload.Phone).Update("token", refreshToken)

	return c.JSON(fiber.Map{
		"routeName":     "login",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
