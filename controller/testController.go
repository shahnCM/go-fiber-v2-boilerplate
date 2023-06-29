package controller

import (
	"boilerplate/service"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	message := service.GetHomeMessage()
	return c.SendString(message)
}
