package service

import (
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/param"
	"bitereview/app/repository"
	"bitereview/app/schema"
	"bitereview/app/serializer"
	"bitereview/app/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstimateService struct {
	estimateRepo *repository.EstimateRepository
	reviewRepo   *repository.ReviewRepository
}

func NewEstimateService(
	estimateRepo *repository.EstimateRepository, reviewRepo *repository.ReviewRepository,
) *EstimateService {
	return &EstimateService{
		estimateRepo: estimateRepo,
		reviewRepo:   reviewRepo,
	}
}

func (es *EstimateService) GetEstimatesByReviewId(c *fiber.Ctx) error {
	reviewId, err := param.ParamPrimitiveID(c, "reviewId")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	limit, offset := param.GetLimitOffset(c)

	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	reviews, err := es.estimateRepo.GetEntitiesByReviewId(timeoutCtx, reviewId, limit, offset)
	if err != nil {
		return helper.SendError(c, err, errors.MakeRepositoryError("Estimate"))
	}

	return helper.SendSuccess(c, reviews)
}

func (es *EstimateService) AddEstimate(c *fiber.Ctx) error {
	reviewId, err := param.ParamPrimitiveID(c, "reviewId")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	data, err := serializer.GetSerializedCreateEstimate(c)
	if err != nil {
		return err
	}

	_, _, err = serializer.GetJwtUserLocalWithParsedID(c)

	if err != nil {
		return helper.SendError(c, err, errors.Unauthorized)
	}

	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	_, err = es.reviewRepo.GetEntityByID(timeoutCtx, reviewId)
	if err == nil {
		return helper.SendError(c, err, errors.EntityNotExists)
	}

	estimate := &schema.Estimate{
		ID:          primitive.NewObjectID(),
		ReviewID:    reviewId,
		Name:        data.Name,
		Value:       data.Value,
		Description: data.Description,
	}

	withTimeout, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	es.estimateRepo.CreateEntity(withTimeout, estimate)

	return helper.SendSuccess(c, estimate)
}
