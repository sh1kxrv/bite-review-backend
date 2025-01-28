package restaurant

import (
	"bitereview/enum"
	"bitereview/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouterRestaurant struct {
	handler *RestaurantHandler
}

func NewRouterRestaurant(service *RestaurantService) *RouterRestaurant {
	return &RouterRestaurant{
		handler: NewRestaurantHandler(service),
	}
}

func (rh *RouterRestaurant) registerModeratorRoutes(g fiber.Router) {
	moderRoute := g.Group("/moderator/restaurant", middleware.JwtAuthMiddleware, middleware.CreateRoleMiddleware(enum.StaffRoles...))
	moderRoute.Patch("/:id/verify", rh.handler.VerifyRestaurant)
	moderRoute.Patch("/:id/unverify", rh.handler.UnverifyRestaurant)
}

func (rh *RouterRestaurant) registerAdminRoutes(g fiber.Router) {
	adminRoute := g.Group("/admin/restaurant", middleware.JwtAuthMiddleware, middleware.CreateRoleMiddleware(enum.RoleAdmin))
	adminRoute.Post("/", rh.handler.CreateRestaurant)
}

func (rh *RouterRestaurant) registerPublicRoutes(g fiber.Router) {
	restRoute := g.Group("/restaurant")
	restRoute.Get("/", rh.handler.GetRestaurants)
	restRoute.Get("/:id", rh.handler.GetRestaurantById)
}

func (rh *RouterRestaurant) RegisterRoutes(g fiber.Router) {
	rh.registerModeratorRoutes(g)
	rh.registerAdminRoutes(g)
	rh.registerPublicRoutes(g)
}
