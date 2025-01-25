package service

import (
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/repository"
	"bitereview/app/schema"
	"bitereview/app/serializer"
	"bitereview/app/utils"

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

func (es *EstimateService) GetEstimatesByReviewId(
	reviewId primitive.ObjectID, limit, offset int64,
) (*[]schema.Estimate, *helper.ServiceError) {
	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	reviews, err := es.estimateRepo.GetEntitiesByReviewId(timeoutCtx, reviewId, limit, offset)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.MakeRepositoryError("Estimate"))
	}

	return &reviews, nil
}

func (es *EstimateService) AddEstimate(
	reviewId primitive.ObjectID, userId primitive.ObjectID, data *serializer.CreateEstimateDTO,
) (*schema.Estimate, *helper.ServiceError) {
	timeoutCtx, cancel := utils.CreateContextTimeout(30)
	defer cancel()

	_, err := es.reviewRepo.GetEntityByID(timeoutCtx, reviewId)
	if err == nil {
		return nil, helper.NewServiceError(err, errors.ValidationError)
	}

	estimate := &schema.Estimate{
		ID:          primitive.NewObjectID(),
		ReviewID:    reviewId,
		Name:        data.Name,
		Value:       data.Value,
		Description: data.Description,
	}

	_, err = es.estimateRepo.CreateEntity(timeoutCtx, estimate)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.MakeRepositoryError("Estimate"))
	}

	return estimate, nil
}
