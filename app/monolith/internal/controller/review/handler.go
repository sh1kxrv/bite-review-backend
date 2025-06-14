package review

import (
	"shared/errors"
	"shared/serializer"
	"shared/utils"
	"shared/utils/helper"
	"shared/utils/param"

	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	reviewService *ReviewService
}

func NewReviewHandler(service *ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: service,
	}
}

// @Summary Создать обзор
// @Tags Обзоры
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param restaurantId path string true "ID ресторана"
// @Success 200 {object} entity.Review
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/public/review/{restaurantId} [post]
func (rh *ReviewHandler) CreateReview(c *fiber.Ctx) error {
	_, userId, err := utils.GetJwtUserLocalWithParsedID(c)

	if err != nil {
		return helper.SendError(c, err, errors.Unauthorized)
	}

	data, err := serializer.GetSerializedReview(c)

	if err != nil {
		return helper.SendError(c, err, errors.RepositoryError)
	}

	restId, err := param.ParamPrimitiveID(c, "restaurantId")

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	review, serr := rh.reviewService.CreateReview(userId, restId, &data)

	return helper.SendSomething(c, &review, serr)
}

// @Summary Получить обзоры
// @Tags Обзоры
// @Accept json
// @Produce json
// @Param restaurantId path string true "ID ресторана"
// @Param limit query int false "Количество"
// @Param offset query int false "Смещение"
// @Success 200 {array} entity.Review
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/public/review/{restaurantId} [get]
func (rh *ReviewHandler) GetReviewsByRestaurantId(c *fiber.Ctx) error {
	restId, err := param.ParamPrimitiveID(c, "restaurantId")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	limit, offset := param.GetLimitOffset(c)

	reviews, serr := rh.reviewService.GetReviewsByRestaurantId(restId, limit, offset)
	return helper.SendSomething(c, reviews, serr)
}
