package handler

import (
	"bitereview/errors"
	"bitereview/helper"
	"bitereview/middleware"
	"bitereview/param"
	"bitereview/serializer"
	"bitereview/service"

	"github.com/gofiber/fiber/v2"
)

type EstimateHandler struct {
	EstimateService *service.EstimateService
}

func NewEstimateHandler(estimateService *service.EstimateService) *EstimateHandler {
	return &EstimateHandler{
		EstimateService: estimateService,
	}
}

// @Summary Получение оценок из ревью
// @Tags Оценка
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param reviewId path string true "ID review"
// @Success 200 {array} schema.Estimate
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/estimate/{reviewId} [get]
func (sh *EstimateHandler) GetEstimatesByReviewId(c *fiber.Ctx) error {
	reviewId, err := param.ParamPrimitiveID(c, "reviewId")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	limit, offset := param.GetLimitOffset(c)

	data, serr := sh.EstimateService.GetEstimatesByReviewId(reviewId, limit, offset)
	return helper.SendSomething(c, &data, serr)
}

// @Summary Добавление оценки в ревью
// @Tags Оценка
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param reviewId path string true "ID review"
// @Param data body serializer.CreateEstimateDTO true "Оценка"
// @Success 200 {object} schema.Estimate
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/estimate/{reviewId} [post]
func (sh *EstimateHandler) AddEstimate(c *fiber.Ctx) error {
	_, userId, err := serializer.GetJwtUserLocalWithParsedID(c)
	if err != nil {
		return helper.SendError(c, err, errors.JwtPairVerificationError)
	}

	reviewId, err := param.ParamPrimitiveID(c, "reviewId")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	payload, err := serializer.GetSerializedCreateEstimate(c)
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	data, serr := sh.EstimateService.AddEstimate(reviewId, userId, &payload)

	return helper.SendSomething(c, &data, serr)
}

func (sh *EstimateHandler) RegisterRoutes(g fiber.Router) {
	estimateRoute := g.Group("/estimate", middleware.JwtAuthMiddleware)

	estimateRoute.Get("/:reviewId", sh.GetEstimatesByReviewId)
	estimateRoute.Post("/:reviewId", sh.AddEstimate)
}
