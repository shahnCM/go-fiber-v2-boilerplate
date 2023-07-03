package controller

import (
	"boilerplate/pkg/helper"
	"boilerplate/pkg/request"
	"github.com/gofiber/fiber/v2"
)

func StoreUser(ctx *fiber.Ctx) error {
	user := new(request.UserStore)

	if err := ctx.BodyParser(user); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := helper.ValidateRequest(user); err != nil {
		return helper.ValidationErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, user, 200)
}
