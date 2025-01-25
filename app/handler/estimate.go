package handler

import (
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/middleware"
	"bitereview/app/param"
	"bitereview/app/serializer"
	"bitereview/app/service"

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

func (sh *EstimateHandler) GetEstimatesByReviewId(c *fiber.Ctx) error {
	reviewId, err := param.ParamPrimitiveID(c, "reviewId")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	limit, offset := param.GetLimitOffset(c)

	data, serr := sh.EstimateService.GetEstimatesByReviewId(reviewId, limit, offset)
	return helper.SendSomething(c, &data, serr)
}

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
