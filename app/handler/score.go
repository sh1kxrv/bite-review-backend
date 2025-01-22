package handler

import (
	"bitereview/app/middleware"
	"bitereview/app/repository"

	"github.com/gofiber/fiber/v2"
)

type ScoreHandler struct {
	ScoreRepo *repository.ScoreRepository
}

func NewScoreHandler(scoreRepo *repository.ScoreRepository) *ScoreHandler {
	return &ScoreHandler{
		ScoreRepo: scoreRepo,
	}
}

func (sh *ScoreHandler) GetScoreByReviewId(c *fiber.Ctx) error {
	return nil
}

func (sh *ScoreHandler) RegisterRoutes(g fiber.Router) {
	scoreRoute := g.Group("/score", middleware.JwtAuthMiddleware)

	scoreRoute.Get("/:reviewId", sh.GetScoreByReviewId)
}
