package review

import (
	"bitereview/utils"

	"github.com/gofiber/fiber/v2"
)

type ReviewDTO struct {
	Summary string `json:"summary" validate:"required,min=64,max=1024"`
}

func GetSerializedReview(c *fiber.Ctx) (ReviewDTO, error) {
	return utils.GetSerializedBodyData[ReviewDTO](c)
}
