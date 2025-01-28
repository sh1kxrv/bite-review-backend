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
	estimatePublicRoute := g.Group("/public/estimate", middleware.JwtAuthMiddleware)
	estimatePrivateGroup := g.Group("/estimate", middleware.JwtAuthMiddleware)

	estimatePublicRoute.Get("/:reviewId", re.handler.GetEstimatesByReviewId)
	estimatePrivateGroup.Post("/:reviewId", re.handler.AddEstimate)
}
