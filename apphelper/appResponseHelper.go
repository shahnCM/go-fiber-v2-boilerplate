package apphelper

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

func SuccessResponse(data interface{}, code int) SuccessResponseBody {
	return SuccessResponseBody{
		Status:     "Success",
		StatusCode: code,
		Data:       data,
	}
}

func ErrorResponse(code int, message string) ErrorResponseBody {
	return ErrorResponseBody{
		Status:       "Error",
		StatusCode:   code,
		ErrorMessage: message,
	}
}

func ValidationErrorResponse(validationErrors []ValidationError) ValidationErrorResponseBody {
	return ValidationErrorResponseBody{
		ErrorResponseBody: ErrorResponseBody{
			Status:       "Error",
			StatusCode:   422,
			ErrorMessage: "Validation Error",
		},
		ValidationErrors: validationErrors,
	}
}
