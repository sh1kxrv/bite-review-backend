package handler

import (
	"bitereview/middleware"
	"bitereview/service"

	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	reviewService *service.ReviewService
}

func NewReviewHandler(service *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: service,
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
