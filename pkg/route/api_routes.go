package route

import (
	"boilerplate/pkg/controller"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(app *fiber.App) {
	app.Get("/home", controller.Home)
	app.Post("/user", controller.StoreUser)
	app.Get("/fiber-error", controller.FiberError)
	app.Get("/simple-panic", controller.SimplePanic)
	app.Get("/custom-error", controller.CustomError)
	app.Get("/null-response", controller.NullResponse)
}
