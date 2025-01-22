package handler

import (
	"bitereview/app/middleware"
	"bitereview/app/repository"

	"github.com/gofiber/fiber/v2"
)

type RestaurantHandler struct {
	RestaurantRepo *repository.RestaurantRepository
}

func NewRestaurantHandler(restaurantRepo *repository.RestaurantRepository) *RestaurantHandler {
	return &RestaurantHandler{
		RestaurantRepo: restaurantRepo,
	}
}

func (rh *RestaurantHandler) GetRestaurants(c *fiber.Ctx) error {
	return nil
}

func (rh *RestaurantHandler) GetRestaurantById(c *fiber.Ctx) error {
	return nil
}

func (rh *RestaurantHandler) RegisterRoutes(g fiber.Router) {
	restRoute := g.Group("/restaurant", middleware.JwtAuthMiddleware)

	// TODO: Limit / Offset
	restRoute.Get("/", rh.GetRestaurants)

	restRoute.Get("/:id", rh.GetRestaurantById)
}
