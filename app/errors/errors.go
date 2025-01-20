package errors

import (
	"bitereview/app/helper"
	"net/http"
)

var UnknownError = helper.ErrorResponse{
	StatusCode: http.StatusInternalServerError,
	Message:    "Unknown error",
}

var CryptoError = helper.ErrorResponse{
	StatusCode: http.StatusInternalServerError,
	Message:    "Crypto error",
}
