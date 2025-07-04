package auth

import (
	"shared/errors"
	"shared/serializer"
	"shared/utils/helper"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// @Summary Авторизация
// @Tags Авторизация
// @Accept json
// @Produce json
// @Param data body dto.AuthDataLogin true "Авторизационные данные"
// @Success 200 {object} ro.JwtPair
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	data, err := serializer.GetSerializedAuthLoginData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	pair, serr := h.authService.Login(data.Email, data.Password)
	return helper.SendSomething(c, &pair, serr)
}

// @Summary Регистрация
// @Tags Авторизация
// @Accept json
// @Produce json
// @Param data body dto.AuthDataRegister true "Данные регистрации"
// @Success 200 {object} ro.JwtPair
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	data, err := serializer.GetSerializedAuthRegisterData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	pair, serr := h.authService.Register(&data)
	return helper.SendSomething(c, &pair, serr)
}

// @Summary Обновление Access токена
// @Tags Авторизация
// @Accept json
// @Produce json
// @Param data body dto.AuthDataRefresh true "Данные обновления токена"
// @Success 200 {object} ro.JwtPair
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/auth/refresh [post]
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
