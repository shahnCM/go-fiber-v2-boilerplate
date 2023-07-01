package apphelper

import "github.com/gofiber/fiber/v2"

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
	ValidationErrors []ValidationError `json:"validation_errors"`
}

type ValidationError struct {
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

func ValidationErrorResponse(ctx *fiber.Ctx, validationErrors *[]ValidationError) error {
	return ctx.Status(422).JSON(ValidationErrorResponseBody{
		ErrorResponseBody: ErrorResponseBody{
			Status:       "Error",
			StatusCode:   422,
			ErrorMessage: "Validation Error",
		},
		ValidationErrors: *validationErrors,
	})
}
