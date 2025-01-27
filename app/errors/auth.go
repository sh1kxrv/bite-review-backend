package errors

import (
	"bitereview/helper"

	"github.com/gofiber/fiber/v2"
)

var Unauthorized = &helper.ErrorResponse{
	StatusCode: fiber.StatusUnauthorized,
	Message:    "Unauthorized",
}

var Forbidden = &helper.ErrorResponse{
	StatusCode: fiber.StatusForbidden,
	Message:    "Forbidden",
}
