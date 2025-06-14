package review

import (
	"bitereview/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouterReview struct {
	handler *ReviewHandler
}

func NewRouterReview(service *ReviewService) *RouterReview {
	return &RouterReview{
		handler: NewReviewHandler(service),
	}
}

func (rr *RouterReview) RegisterRoutes(g fiber.Router) {
	reviewGroupPrivate := g.Group("/review", middleware.JwtAuthMiddleware)
	reviewGroupPublic := g.Group("/public/review")

	reviewGroupPrivate.Post("/:restaurantId", rr.handler.CreateReview)
	reviewGroupPublic.Get("/:restaurantId", rr.handler.GetReviewsByRestaurantId)
}
