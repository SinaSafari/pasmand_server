package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"pasmand/internal/config"
	"pasmand/internal/routes"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//go kafka.Consume()
	app := fiber.New()
	config.SetupDependencies()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
