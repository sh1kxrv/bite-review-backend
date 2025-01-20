package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func SendError(c *fiber.Ctx, origErr error, e ErrorResponse) error {
	resp := Response{
		Status: false,
		Data:   e,
	}
	if origErr != nil {
		logrus.Errorf("Fictive error %s, real error: %s", e.Message, origErr.Error())
	}
	return c.Status(e.StatusCode).JSON(resp)
}

func buildSuccessResponse(data any) Response {
	resp := Response{
		Status: true,
		Data:   data,
	}
	return resp
}

func SendSuccess(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(buildSuccessResponse(data))
}
