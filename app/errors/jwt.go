package errors

import (
	"bitereview/app/helper"
	"net/http"
)

var JwtPairGenerationError = helper.ErrorResponse{
	StatusCode: http.StatusInternalServerError,
	Message:    "Could not generate JWT pair",
}

var JwtPairVerificationError = helper.ErrorResponse{
	StatusCode: http.StatusUnauthorized,
	Message:    "Could not verify JWT pair",
}

var JwtRefreshTokenInvalid = helper.ErrorResponse{
	StatusCode: http.StatusUnauthorized,
	Message:    "Invalid JWT refresh token",
}
