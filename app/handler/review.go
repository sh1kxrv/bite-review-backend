package handler

import (
	"bitereview/app/middleware"
	"bitereview/app/repository"

	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	ReviewRepo *repository.ReviewRepository
}

func NewReviewHandler(reviewRepo *repository.ReviewRepository) *ReviewHandler {
	return &ReviewHandler{
		ReviewRepo: reviewRepo,
	}
}

func (rh *ReviewHandler) GetReviewsByRestaurantId(c *fiber.Ctx) error {
	return nil
}

func (rh *ReviewHandler) RegisterRoutes(g fiber.Router) {
	reviewRoute := g.Group("/review", middleware.JwtAuthMiddleware)

	// TODO: Limit / Offset
	reviewRoute.Get("/:restaurantId", rh.GetReviewsByRestaurantId)
}
