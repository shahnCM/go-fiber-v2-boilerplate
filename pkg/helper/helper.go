package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SuccessResponseBody struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type ErrorResponseBody struct {
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}

type ValidationErrorResponseBody struct {
	ErrorResponseBody
	ValidationErrors []*ValidationErrorElement `json:"validation_errors"`
}

type ValidationErrorElement struct {
	Msg         string `json:"msg"`
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func SuccessResponse(ctx *fiber.Ctx, data interface{}, code int) error {
	return ctx.Status(code).JSON(SuccessResponseBody{
		Status:     "Success",
		StatusCode: code,
		Data:       data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(ErrorResponseBody{
		Status:       "Error",
		StatusCode:   code,
		ErrorMessage: message,
	})
}

func ValidationErrorResponse(ctx *fiber.Ctx, validationErrorElements []*ValidationErrorElement) error {
	return ctx.Status(422).JSON(ValidationErrorResponseBody{
		ErrorResponseBody: ErrorResponseBody{
			Status:       "Error",
			StatusCode:   422,
			ErrorMessage: "Validation Error",
		},
		ValidationErrors: validationErrorElements,
	})
}

func ParseRequest(ctx *fiber.Ctx, input interface{}) error {
	if err := ctx.BodyParser(input); err != nil {
		return fiber.ErrInternalServerError
	}
	return nil
}

func ValidateRequest(input interface{}) []*ValidationErrorElement {
	var validate = validator.New()
	var errors []*ValidationErrorElement
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErrorElement
			value := err.Param()
			if value == "" {
				value = "missing or invalid"
			}
			element.Msg = fmt.Sprintf(`%s, %s: %s`, err.Field(), err.Tag(), value)
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
