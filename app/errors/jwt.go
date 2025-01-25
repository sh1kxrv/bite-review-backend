package errors

import (
	"bitereview/app/helper"

	"github.com/gofiber/fiber/v2"
)

var JwtPairGenerationError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "Could not generate JWT pair",
}

var JwtPairVerificationError = &helper.ErrorResponse{
	StatusCode: fiber.StatusUnauthorized,
	Message:    "Could not verify JWT pair",
}

var JwtRefreshTokenInvalid = &helper.ErrorResponse{
	StatusCode: fiber.StatusUnauthorized,
	Message:    "Invalid JWT refresh token",
}
