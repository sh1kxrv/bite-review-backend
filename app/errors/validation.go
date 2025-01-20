package errors

import (
	"bitereview/app/helper"
	"net/http"
)

var ValidationError = helper.ErrorResponse{
	StatusCode: http.StatusBadRequest,
	Message:    "Validation error",
}

var ParseIDError = helper.ErrorResponse{
	StatusCode: http.StatusBadRequest,
	Message:    "Parse ID error",
}

func MakeValidationError(err error) helper.ErrorResponse {
	return helper.ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
	}
}
