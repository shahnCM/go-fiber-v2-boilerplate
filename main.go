package main

import (
	"boilerplate/pkg/errorhandler"
	"boilerplate/pkg/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	// Configs
	fiberConfig := fiber.Config{ErrorHandler: errorhandler.CustomFiberErrorHandler}
	recoverConfig := recover.Config{EnableStackTrace: true}
	// New app
	app := fiber.New(fiberConfig)
	// Recover from any panic
	app.Use(recover.New(recoverConfig))
	// Initialize api route
	route.InitApiRoutes(app)
	// Start the server
	log.Fatal(app.Listen(":3000"))
}
