package errors

import (
	"bitereview/app/helper"

	"github.com/gofiber/fiber/v2"
)

var ValidationError = &helper.ErrorResponse{
	StatusCode: fiber.StatusBadRequest,
	Message:    "Validation error",
}

var ParseIDError = &helper.ErrorResponse{
	StatusCode: fiber.StatusBadRequest,
	Message:    "Parse ID error",
}

func MakeValidationError(err error) *helper.ErrorResponse {
	return &helper.ErrorResponse{
		StatusCode: fiber.StatusBadRequest,
		Message:    err.Error(),
	}
}
