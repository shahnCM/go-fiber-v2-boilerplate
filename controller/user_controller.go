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

	if validationErrors := apphelper.ValidateRequest(*user); validationErrors != nil {
		return apphelper.ValidationErrorResponse(ctx, validationErrors)
	}

	return apphelper.SuccessResponse(ctx, user, 200)
}
