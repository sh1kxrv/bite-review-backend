package serializer

import (
	"shared/transfer/dto"
	"shared/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSerializedCreateEstimate(c *fiber.Ctx) (dto.CreateEstimateDTO, error) {
	return utils.GetSerializedBodyData[dto.CreateEstimateDTO](c)
}
