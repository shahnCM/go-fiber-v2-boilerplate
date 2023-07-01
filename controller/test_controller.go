package controller

import (
	"boilerplate/apphelper"
	"boilerplate/service"

	"github.com/gofiber/fiber/v2"
)

func Home(ctx *fiber.Ctx) error {
	message := service.GetHomeMessage()
	data := map[string]interface{}{
		"msg": message,
	}

	return ctx.Status(200).JSON(apphelper.SuccessResponse(data, 200))
}

func NullResponse(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(apphelper.SuccessResponse(nil, 201))
}

func SimplePanic(ctx *fiber.Ctx) error {
	if 2 == 2 {
		panic("This panic is caught by fiber")
	}

	return ctx.SendString("Recovery Test")
}

func FiberError(ctx *fiber.Ctx) error {
	// 503 Service Unavailable
	return fiber.ErrServiceUnavailable
}

func CustomError(ctx *fiber.Ctx) error {
	// 503 Service Unavailable
	return fiber.NewError(fiber.StatusServiceUnavailable, "My Custom Error")
}
