package review

import (
	"bitereview/middleware"

	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	reviewService *ReviewService
}

func NewReviewHandler(service *ReviewService) *ReviewHandler {
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
