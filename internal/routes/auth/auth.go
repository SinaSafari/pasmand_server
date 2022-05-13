package auth

import (
	"github.com/gofiber/fiber/v2"
	"pasmand/internal/handlers/auth"
	"pasmand/internal/middlewares"
)

func SetupAuthRoutes(app fiber.Router) {
	authRouter := app.Group("/auth")
	authRouter.Post("/login-request", auth.LoginRequest)
	authRouter.Post("/login", auth.Login)
	authRouter.Post("/refresh", auth.Refresh)

	authRouter.Use(middlewares.AuthMiddleware).Post("/logout", auth.Logout)
}
