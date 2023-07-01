package apphelper

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
)

func ParseRequest(ctx *fiber.Ctx, input interface{}) error {
	if err := ctx.BodyParser(input); err != nil {
		return fiber.ErrInternalServerError
	}
	return nil
}

func ValidateRequest(input interface{}) *[]ValidationError {
	var validate = validator.New()
	var errors []ValidationError
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
	}
	return &errors
}
