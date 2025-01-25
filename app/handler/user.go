package handler

import (
	"bitereview/app/helper"
	"bitereview/app/middleware"
	"bitereview/app/serializer"
	"bitereview/app/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (h *UserHandler) GetMeHandler(c *fiber.Ctx) error {
	_, parsedId, err := serializer.GetJwtUserLocalWithParsedID(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	user, serr := h.UserService.GetUserByID(parsedId)
	return helper.SendSomething(c, &user, serr)
}

func (h *UserHandler) RegisterRoutes(g fiber.Router) {
	userRoute := g.Group("/user", middleware.JwtAuthMiddleware)
	userRoute.Get("/me", h.GetMeHandler)
}
