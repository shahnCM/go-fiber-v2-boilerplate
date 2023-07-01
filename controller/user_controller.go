package controller

import (
	"boilerplate/apphelper"
	"boilerplate/requestvalidation"

	"github.com/gofiber/fiber/v2"
)

func StoreUser(ctx *fiber.Ctx) error {
	user := new(requestvalidation.User)

	if err := ctx.BodyParser(user); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := apphelper.ValidateRequest(user); *err != nil {
		return apphelper.ValidationErrorResponse(ctx, err)
	}

	return apphelper.SuccessResponse(ctx, user, 200)
}
