package estimate

import (
	"shared/database/mongodb/entity"
	"shared/database/mongodb/repository"
	"shared/errors"
	"shared/transfer/dto"
	"shared/utils"
	"shared/utils/helper"

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
) (*[]entity.Estimate, *helper.ServiceError) {
	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	reviews, err := es.estimateRepo.GetEntitiesByReviewId(timeoutCtx, reviewId, limit, offset)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.MakeRepositoryError("Estimate"))
	}

	return &reviews, nil
}

func (es *EstimateService) AddEstimate(
	reviewId primitive.ObjectID, userId primitive.ObjectID, data *dto.CreateEstimateDTO,
) (*entity.Estimate, *helper.ServiceError) {
	timeoutCtx, cancel := utils.CreateContextTimeout(30)
	defer cancel()

	_, err := es.reviewRepo.GetEntityByID(timeoutCtx, reviewId)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.ValidationError)
	}

	estimate := &entity.Estimate{
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
