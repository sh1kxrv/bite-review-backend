package user

import (
	"bitereview/helper"
	"bitereview/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

// @Summary Получить данные о своём аккаунте
// @Tags Пользователь
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.User
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/user/me [get]
func (h *UserHandler) GetMeHandler(c *fiber.Ctx) error {
	_, parsedId, err := utils.GetJwtUserLocalWithParsedID(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	user, serr := h.UserService.GetUserByID(parsedId)
	return helper.SendSomething(c, &user, serr)
}
