package estimate

import (
	"bitereview/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouterEstimate struct {
	handler *EstimateHandler
}

func NewRouterEstimate(service *EstimateService) *RouterEstimate {
	return &RouterEstimate{
		handler: NewEstimateHandler(service),
	}
}

func (re *RouterEstimate) RegisterRoutes(g fiber.Router) {
	estimateRoute := g.Group("/estimate", middleware.JwtAuthMiddleware)

	estimateRoute.Get("/:reviewId", re.handler.GetEstimatesByReviewId)
	estimateRoute.Post("/:reviewId", re.handler.AddEstimate)
}
