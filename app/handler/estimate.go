package handler

import (
	"bitereview/app/middleware"
	"bitereview/app/service"

	"github.com/gofiber/fiber/v2"
)

type EstimateHandler struct {
	EstimateService *service.EstimateService
}

func NewEstimateHandler(estimateService *service.EstimateService) *EstimateHandler {
	return &EstimateHandler{
		EstimateService: estimateService,
	}
}

func (sh *EstimateHandler) GetEstimatesByReviewId(c *fiber.Ctx) error {
	return sh.EstimateService.GetEstimatesByReviewId(c)
}

// TODO
func (sh *EstimateHandler) AddEstimate(c *fiber.Ctx) error {
	return nil
}

func (sh *EstimateHandler) RegisterRoutes(g fiber.Router) {
	estimateRoute := g.Group("/estimate", middleware.JwtAuthMiddleware)

	estimateRoute.Get("/:reviewId", sh.GetEstimatesByReviewId)
	estimateRoute.Post("/:reviewId", sh.AddEstimate)
}
