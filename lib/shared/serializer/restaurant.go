package serializer

import (
	"shared/transfer/dto"
	"shared/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSerializedCreateRestaurant(c *fiber.Ctx) (dto.CreateRestaurantDTO, error) {
	return utils.GetSerializedBodyData[dto.CreateRestaurantDTO](c)
}
