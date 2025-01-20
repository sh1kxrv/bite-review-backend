package errors

import (
	"bitereview/app/helper"
	"net/http"
)

func MakeRepositoryError(repository string) helper.ErrorResponse {
	return helper.ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Repository error: " + repository,
	}
}

var EntityAlreadyExists = helper.ErrorResponse{
	StatusCode: http.StatusConflict,
	Message:    "Entity already exists",
}

var EntityNotExists = helper.ErrorResponse{
	StatusCode: http.StatusNotFound,
	Message:    "Entity not exists",
}

var RepositoryError = helper.ErrorResponse{
	StatusCode: http.StatusInternalServerError,
	Message:    "Repository error",
}
