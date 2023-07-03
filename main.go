package main

import (
	"boilerplate/pkg/config"
	"boilerplate/pkg/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"os"
)

func main() {
	// New app
	app := fiber.New(config.FiberConfig())
	// Recover from any panic
	app.Use(recover.New(config.RecoveryConfig()))
	// Initialize api route
	route.InitApiRoutes(app)
	// Start the server
	if port := os.Getenv("PORT"); port == "" {
		port = "8000"
		log.Fatal(app.Listen(":" + port))
	}

}
