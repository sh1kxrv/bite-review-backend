package review

import (
	"bitereview/middleware"

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
	g.Group("/review", middleware.JwtAuthMiddleware)
}
