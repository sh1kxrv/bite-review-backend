package handler

import (
	"bitereview/enum"
	"bitereview/errors"
	"bitereview/helper"
	"bitereview/middleware"
	"bitereview/param"
	"bitereview/serializer"
	"bitereview/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RestaurantHandler struct {
	RestaurantService *service.RestaurantService
}

func NewRestaurantHandler(service *service.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{
		RestaurantService: service,
	}
}

// @Summary Получить рестораны
// @Tags Рестораны | Public
// @Accept json
// @Produce json
// @Param limit query int false "Количество садов"
// @Param offset query int false "Смещение по количеству"
// @Success 200 {array} schema.Restaurant
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/restaurant [get]
func (rh *RestaurantHandler) GetRestaurants(c *fiber.Ctx) error {
	limit, offset := param.GetLimitOffset(c)

	logrus.Debugf("Get restaurants limit: %d, offset: %d", limit, offset)

	restaurants, serr := rh.RestaurantService.GetRestaurants(limit, offset)
	return helper.SendSomething(c, &restaurants, serr)
}

// @Summary Получить ресторан по ID
// @Tags Рестораны | Public
// @Accept json
// @Produce json
// @Param id path string true "ID ресторана"
// @Success 200 {object} schema.Restaurant
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/restaurant/{id} [get]
func (rh *RestaurantHandler) GetRestaurantById(c *fiber.Ctx) error {
	id, err := param.ParamPrimitiveID(c, "id")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}
	restaurant, serr := rh.RestaurantService.GetRestaurantById(id)
	return helper.SendSomething(c, &restaurant, serr)
}

// @Summary Регистрация ресторана в системе
// @Tags Рестораны | Admin
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body serializer.CreateRestaurantDTO true "Данные регистрации ресторана"
// @Success 200 {object} schema.Restaurant
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/admin/restaurant [post]
func (rh *RestaurantHandler) CreateRestaurant(c *fiber.Ctx) error {
	data, err := serializer.GetSerializedCreateRestaurant(c)
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}
	restaurant, serr := rh.RestaurantService.CreateRestaurant(&data)
	return helper.SendSomething(c, &restaurant, serr)
}

func (rh *RestaurantHandler) setVerifyState(c *fiber.Ctx, verifyBool bool) error {
	id, err := param.ParamPrimitiveID(c, "id")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}
	serr := rh.RestaurantService.UpdateVerifiedStatus(id, verifyBool)
	return helper.SendSomething(c, nil, serr)
}

// @Summary Верификция ресторана
// @Tags Рестораны | Moderator
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "ID ресторана"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/moderator/restaurant/{id}/verify [patch]
func (rh *RestaurantHandler) VerifyRestaurant(c *fiber.Ctx) error {
	return rh.setVerifyState(c, true)
}

// @Summary Отмена верификации ресторана
// @Tags Рестораны | Moderator
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "ID ресторана"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/moderator/restaurant/{id}/unverify [patch]
func (rh *RestaurantHandler) UnverifyRestaurant(c *fiber.Ctx) error {
	return rh.setVerifyState(c, false)
}

func (rh *RestaurantHandler) registerModeratorRoutes(g fiber.Router) {
	moderRoute := g.Group("/moderator/restaurant", middleware.JwtAuthMiddleware, middleware.CreateRoleMiddleware(enum.StaffRoles...))
	moderRoute.Patch("/:id/verify", rh.VerifyRestaurant)
	moderRoute.Patch("/:id/unverify", rh.UnverifyRestaurant)
}

func (rh *RestaurantHandler) registerAdminRoutes(g fiber.Router) {
	adminRoute := g.Group("/admin/restaurant", middleware.JwtAuthMiddleware, middleware.CreateRoleMiddleware(enum.RoleAdmin))
	adminRoute.Post("/", rh.CreateRestaurant)
}

func (rh *RestaurantHandler) registerPublicRoutes(g fiber.Router) {
	restRoute := g.Group("/restaurant")
	restRoute.Get("/", rh.GetRestaurants)
	restRoute.Get("/:id", rh.GetRestaurantById)
}

func (rh *RestaurantHandler) RegisterRoutes(g fiber.Router) {
	rh.registerModeratorRoutes(g)
	rh.registerAdminRoutes(g)
	rh.registerPublicRoutes(g)
}
