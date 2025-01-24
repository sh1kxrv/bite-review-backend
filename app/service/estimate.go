package service

import (
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/param"
	"bitereview/app/repository"
	"bitereview/app/utils"

	"github.com/gofiber/fiber/v2"
)

type EstimateService struct {
	EstimateRepo *repository.EstimateRepository
}

func NewEstimateService(estimateRepo *repository.EstimateRepository) *EstimateService {
	return &EstimateService{
		EstimateRepo: estimateRepo,
	}
}

func (es *EstimateService) GetEstimatesByReviewId(c *fiber.Ctx) error {
	reviewId, err := utils.ParamPrimitiveID(c, "reviewId")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	limit, offset := param.GetLimitOffset(c)

	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	reviews, err := es.EstimateRepo.GetEntitiesByReviewId(timeoutCtx, reviewId, limit, offset)
	if err != nil {
		return helper.SendError(c, err, errors.MakeRepositoryError("Estimate"))
	}

	return helper.SendSuccess(c, reviews)
}
