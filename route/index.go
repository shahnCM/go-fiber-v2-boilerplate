package route

import (
	"boilerplate/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Post("/", controller.Home)
	// Add more routes as needed
}
