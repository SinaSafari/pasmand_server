package categories

import (
	"github.com/gofiber/fiber/v2"
	"pasmand/internal/handlers/categories"
)

func SetupCategoriesRoutes(app fiber.Router) {
	categoriesRouter := app.Group("/categories")
	categoriesRouter.Get("/", categories.GetAllCategories)
	categoriesRouter.Get("/:id", categories.GetOneCategory)
	categoriesRouter.Post("/", categories.CreateCategory)
	categoriesRouter.Put("/:id", categories.UpdateCategory)
	categoriesRouter.Delete("/:id", categories.DeleteCategory)
}
