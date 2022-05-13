package addresses

import (
	"github.com/gofiber/fiber/v2"
	"pasmand/internal/handlers/address"
)

func SetupAddressesRoutes(app fiber.Router) {
	addressesRouter := app.Group("/addresses")
	addressesRouter.Get("/", address.GetAllAddresses)
	addressesRouter.Get("/:id", address.GetOneAddress)
	addressesRouter.Post("/", address.CreateAddress)
	addressesRouter.Put("/:id", address.UpdateAddress)
	addressesRouter.Delete("/:id", address.DeleteAddress)
}
