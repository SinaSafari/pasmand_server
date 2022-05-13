package routes

import (
	"github.com/gofiber/fiber/v2"
	"pasmand/internal/routes/addresses"
	"pasmand/internal/routes/auth"
	"pasmand/internal/routes/categories"
	"pasmand/internal/routes/products"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	apiV1 := api.Group("/v1")

	auth.SetupAuthRoutes(apiV1)
	products.SetupProductsRoutes(apiV1)
	categories.SetupCategoriesRoutes(apiV1)
	addresses.SetupAddressesRoutes(apiV1)
}
