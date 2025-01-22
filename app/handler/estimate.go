package handler

import (
	"bitereview/app/middleware"
	"bitereview/app/repository"

	"github.com/gofiber/fiber/v2"
)

type EstimateHandler struct {
	EstimateRepo *repository.EstimateRepository
}

func NewEstimateHandler(estimateRepo *repository.EstimateRepository) *EstimateHandler {
	return &EstimateHandler{
		EstimateRepo: estimateRepo,
	}
}

// TODO
func (sh *EstimateHandler) GetEstimatesByReviewId(c *fiber.Ctx) error {
	return nil
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
