package estimate

import (
	"bitereview/errors"
	"bitereview/helper"
	"bitereview/param"
	"bitereview/utils"

	"github.com/gofiber/fiber/v2"
)

type EstimateHandler struct {
	EstimateService *EstimateService
}

func NewEstimateHandler(estimateService *EstimateService) *EstimateHandler {
	return &EstimateHandler{
		EstimateService: estimateService,
	}
}

// @Summary Получение оценок из обзора
// @Tags Оценка,Оценка / Общедоступные
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param reviewId path string true "ID review"
// @Param limit query int false "Количество"
// @Param offset query int false "Смещение"
// @Success 200 {array} entity.Estimate
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/public/estimate/{reviewId} [get]
func (sh *EstimateHandler) GetEstimatesByReviewId(c *fiber.Ctx) error {
	reviewId, err := param.ParamPrimitiveID(c, "reviewId")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	limit, offset := param.GetLimitOffset(c)

	data, serr := sh.EstimateService.GetEstimatesByReviewId(reviewId, limit, offset)
	return helper.SendSomething(c, &data, serr)
}

// @Summary Добавление оценки в обзор
// @Tags Оценка
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param reviewId path string true "ID review"
// @Param data body CreateEstimateDTO true "Оценка"
// @Success 200 {object} entity.Estimate
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/estimate/{reviewId} [post]
func (sh *EstimateHandler) AddEstimate(c *fiber.Ctx) error {
	_, userId, err := utils.GetJwtUserLocalWithParsedID(c)
	if err != nil {
		return helper.SendError(c, err, errors.JwtPairVerificationError)
	}

	reviewId, err := param.ParamPrimitiveID(c, "reviewId")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	payload, err := GetSerializedCreateEstimate(c)
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	data, serr := sh.EstimateService.AddEstimate(reviewId, userId, &payload)

	return helper.SendSomething(c, &data, serr)
}
