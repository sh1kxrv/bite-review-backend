package serializer

import (
	"shared/transfer/dto"
	"shared/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSerializedReview(c *fiber.Ctx) (dto.ReviewDTO, error) {
	return utils.GetSerializedBodyData[dto.ReviewDTO](c)
}
