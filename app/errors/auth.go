package errors

import (
	"bitereview/app/helper"
	"net/http"
)

var Unauthorized = helper.ErrorResponse{
	StatusCode: http.StatusUnauthorized,
	Message:    "Unauthorized",
}

var Forbidden = helper.ErrorResponse{
	StatusCode: http.StatusForbidden,
	Message:    "Forbidden",
}
