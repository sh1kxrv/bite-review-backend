package handler

import (
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/serializer"
	"bitereview/app/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	data, err := serializer.GetSerializedAuthLoginData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	pair, serr := h.authService.Login(data.Email, data.Password)
	return helper.SendSomething(c, &pair, serr)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	data, err := serializer.GetSerializedAuthRegisterData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	pair, serr := h.authService.Register(&data)
	return helper.SendSomething(c, &pair, serr)
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	v, err := serializer.GetSerializedAuthRefreshData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	pair, serr := h.authService.Refresh(v.RefreshToken)
	return helper.SendSomething(c, &pair, serr)
}

func (h *AuthHandler) RegisterRoutes(g fiber.Router) {
	authRoute := g.Group("/auth")

	authRoute.Post("/login", h.Login)
	authRoute.Post("/register", h.Register)
	authRoute.Post("/refresh", h.Refresh)
}
