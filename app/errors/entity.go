package errors

import (
	"bitereview/helper"

	"github.com/gofiber/fiber/v2"
)

func MakeRepositoryError(repository string) *helper.ErrorResponse {
	return &helper.ErrorResponse{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "Repository error: " + repository,
	}
}

var EntityAlreadyExists = &helper.ErrorResponse{
	StatusCode: fiber.StatusConflict,
	Message:    "Entity already exists",
}

var EntityNotExists = &helper.ErrorResponse{
	StatusCode: fiber.StatusNotFound,
	Message:    "Entity not exists",
}

var RepositoryError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "Repository error",
}
