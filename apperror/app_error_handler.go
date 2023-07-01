package apperror

import (
	"errors"

	"boilerplate/apphelper"

	"github.com/gofiber/fiber/v2"
)

func CustomFiberErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := 500
	message := "Internal Server Error"

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error

	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	return apphelper.ErrorResponse(ctx, code, message)
}
