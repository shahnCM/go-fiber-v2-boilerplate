package main

import (
	"log"

	"boilerplate/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Initialize routes
	route.InitRoutes(app)
	// Start the server
	err := app.Listen(":3000")
	// Check for Error
	if err != nil {
		log.Fatal(err)
	}
}
