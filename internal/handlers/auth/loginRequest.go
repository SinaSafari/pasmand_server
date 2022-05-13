package auth

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
	"math/rand"
	"net/http"
	"pasmand/internal/config/database"
	"pasmand/internal/config/kafka"
	"pasmand/internal/config/redis"
	"pasmand/internal/models"
	"strconv"
	"time"
)

type LoginRequestPayload struct {
	Phone string `json:"phone"`
}

// LoginRequest sends otp email/sms and store the code in cache
func LoginRequest(c *fiber.Ctx) error {
	ctx := context.Background()
	payload := new(LoginRequestPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "body is not in correct format",
		})
	}

	database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.User{
		Phone: payload.Phone,
	})

	otpValue := strconv.Itoa(rangeIn(1000, 9999))

	err := redis.RedisClient.Set(ctx, payload.Phone, otpValue, time.Second*180).Err()
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(fiber.Map{
			"status": "redis server not available",
			"data":   err,
		})
	}

	kafka.WriteToKafka(payload.Phone + "," + otpValue)

	return c.JSON(fiber.Map{
		"routeName": "login request",
		"data": fiber.Map{
			"phone":    payload.Phone,
			"otpValue": otpValue,
		},
	})
}

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
