package estimate

import (
	"bitereview/utils"

	"github.com/gofiber/fiber/v2"
)

type CreateEstimateDTO struct {
	Name        string `json:"name" validate:"required,min=2,max=50"`
	Description string `json:"description" validate:"required,min=2,max=50"`
	Value       int    `json:"value" validate:"required,min=1,max=100"`
}

func GetSerializedCreateEstimate(c *fiber.Ctx) (CreateEstimateDTO, error) {
	return utils.GetSerializedBodyData[CreateEstimateDTO](c)
}
