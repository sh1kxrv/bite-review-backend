package user

import (
	"bitereview/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouterUser struct {
	handler *UserHandler
}

func NewRouterUser(service *UserService) *RouterUser {
	return &RouterUser{
		handler: NewUserHandler(service),
	}
}

func (ru *RouterUser) RegisterRoutes(g fiber.Router) {
	userRoute := g.Group("/user", middleware.JwtAuthMiddleware)

	userRoute.Get("/me", ru.handler.GetMeHandler)
}
