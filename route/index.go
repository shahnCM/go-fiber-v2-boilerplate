package route

import (
	"boilerplate/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Post("user", controller.StoreUser)
	app.Get("home", controller.Home)
	app.Get("null-response", controller.NullResponse)
	app.Get("simple-panic", controller.SimplePanic)
	app.Get("fiber-error", controller.FiberError)
	app.Get("custom-error", controller.CustomError)
	// Add more routes as needed
}
