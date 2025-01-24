package handler

import (
	"bitereview/app/enum"
	"bitereview/app/middleware"
	"bitereview/app/service"

	"github.com/gofiber/fiber/v2"
)

type RestaurantHandler struct {
	RestaurantService *service.RestaurantService
}

func NewRestaurantHandler(service *service.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{
		RestaurantService: service,
	}
}

func (rh *RestaurantHandler) GetRestaurants(c *fiber.Ctx) error {
	return rh.RestaurantService.GetRestaurants(c)
}

func (rh *RestaurantHandler) GetRestaurantById(c *fiber.Ctx) error {
	return rh.RestaurantService.GetRestaurantById(c)
}

func (rh *RestaurantHandler) CreateRestaurant(c *fiber.Ctx) error {
	return rh.RestaurantService.CreateRestaurant(c)
}

// TODO
func (rh *RestaurantHandler) VerifyRestaurant(c *fiber.Ctx) error {
	return nil
}

// TODO
func (rh *RestaurantHandler) UnverifyRestaurant(c *fiber.Ctx) error {
	return nil
}

func (rh *RestaurantHandler) RegisterRoutes(g fiber.Router) {
	adminRoute := g.Group("/admin/restaurant", middleware.JwtAuthMiddleware, middleware.CreateRoleMiddleware(enum.RoleAdmin))
	adminRoute.Post("/", rh.CreateRestaurant)

	moderRoute := g.Group("/moderator/restaurant", middleware.JwtAuthMiddleware, middleware.CreateRoleMiddleware(enum.RoleModerator))
	moderRoute.Patch("/:id/verify", rh.VerifyRestaurant)
	moderRoute.Patch("/:id/unverify", rh.UnverifyRestaurant)

	// Public
	restRoute := g.Group("/restaurant")

	restRoute.Get("/", rh.GetRestaurants)
	restRoute.Get("/:id", rh.GetRestaurantById)
}
