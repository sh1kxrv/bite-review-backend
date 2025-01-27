package router

import (
	"bitereview/handler"

	"github.com/gofiber/fiber/v2"
)

type AppRouter struct {
	UserHandler       *handler.UserHandler
	AuthHandler       *handler.AuthHandler
	RestaurantHandler *handler.RestaurantHandler
	ReviewHandler     *handler.ReviewHandler
	EstimateHandler   *handler.EstimateHandler
}

func NewAppRouter(
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	restaurantHandler *handler.RestaurantHandler,
	reviewHandler *handler.ReviewHandler,
	EstimateHandler *handler.EstimateHandler,
) *AppRouter {
	return &AppRouter{
		UserHandler:       userHandler,
		AuthHandler:       authHandler,
		RestaurantHandler: restaurantHandler,
		ReviewHandler:     reviewHandler,
		EstimateHandler:   EstimateHandler,
	}
}

func (r *AppRouter) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	r.AuthHandler.RegisterRoutes(v1)
	r.UserHandler.RegisterRoutes(v1)
	r.RestaurantHandler.RegisterRoutes(v1)
	r.ReviewHandler.RegisterRoutes(v1)
	r.EstimateHandler.RegisterRoutes(v1)
}
