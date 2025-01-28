package review

import (
	"bitereview/entity"
	"bitereview/errors"
	"bitereview/helper"
	"bitereview/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewService struct {
	reviewRepo *ReviewRepository
}

func NewReviewService(reviewRepo *ReviewRepository) *ReviewService {
	return &ReviewService{
		reviewRepo: reviewRepo,
	}
}

func (s *ReviewService) CreateReview(userId, restId primitive.ObjectID, data *ReviewDTO) (*entity.Review, *helper.ServiceError) {
	review := entity.Review{
		ID:           primitive.NewObjectID(),
		RestaurantID: restId,
		UserID:       userId,
		Summary:      data.Summary,
		CreatedAt:    time.Now(),
	}

	ctx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	_, err := s.reviewRepo.CreateEntity(ctx, &review)
	return &review, helper.NewServiceError(err, errors.MakeRepositoryError("Review"))
}

func (s *ReviewService) GetReviewsByRestaurantId(restId primitive.ObjectID, limit, offset int64) ([]entity.Review, *helper.ServiceError) {
	ctx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	reviews, err := s.reviewRepo.GetAll(ctx, bson.M{"restaurantId": restId}, limit, offset)
	return reviews, helper.NewServiceError(err, errors.MakeRepositoryError("Review"))
}
