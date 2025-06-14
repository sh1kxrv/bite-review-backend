package errors

import (
	"shared/utils/helper"

	"github.com/gofiber/fiber/v2"
)

var UnknownError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "Unknown error",
}

var CryptoError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "Crypto error",
}
