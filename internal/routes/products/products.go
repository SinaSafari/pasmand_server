package products

import (
	"github.com/gofiber/fiber/v2"
	"pasmand/internal/handlers/products"
	"pasmand/internal/middlewares"
)

func SetupProductsRoutes(app fiber.Router) {
	productsRouter := app.Group("/products").Use(middlewares.AuthMiddleware)
	productsRouter.Get("/", products.GetAllProducts)
}
