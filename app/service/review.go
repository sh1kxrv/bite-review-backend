package service

import "bitereview/app/repository"

type ReviewService struct {
	ReviewRepo *repository.ReviewRepository
}

func NewReviewService(reviewRepo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{
		ReviewRepo: reviewRepo,
	}
}
