package router

import (
	"bitereview/app/handler"

	"github.com/gofiber/fiber/v2"
)

type AppRouter struct {
	UserHandler *handler.UserHandler
	AuthHandler *handler.AuthHandler
}

func NewAppRouter(
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
) *AppRouter {
	return &AppRouter{
		UserHandler: userHandler,
		AuthHandler: authHandler,
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
}
