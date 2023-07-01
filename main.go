package main

import (
	"log"

	"boilerplate/apperror"
	"boilerplate/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fiberConfig := fiber.Config{ErrorHandler: apperror.CustomFiberErrorHandler}
	// New app
	app := fiber.New(fiberConfig)
	// Recover from panic
	app.Use(recover.New())
	// Initialize routes
	route.InitRoutes(app)
	// Start the server
	log.Fatal(app.Listen(":3000"))
}
