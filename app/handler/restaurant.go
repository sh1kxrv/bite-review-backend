package handler

import (
	"bitereview/app/enum"
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/middleware"
	"bitereview/app/param"
	"bitereview/app/serializer"
	"bitereview/app/service"

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

func (rh *RestaurantHandler) GetRestaurants(c *fiber.Ctx) error {
	limit, offset := param.GetLimitOffset(c)

	logrus.Debugf("Get restaurants limit: %d, offset: %d", limit, offset)

	restaurants, serr := rh.RestaurantService.GetRestaurants(limit, offset)
	return helper.SendSomething(c, &restaurants, serr)
}

func (rh *RestaurantHandler) GetRestaurantById(c *fiber.Ctx) error {
	id, err := param.ParamPrimitiveID(c, "id")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}
	restaurant, serr := rh.RestaurantService.GetRestaurantById(id)
	return helper.SendSomething(c, &restaurant, serr)
}

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

func (rh *RestaurantHandler) VerifyRestaurant(c *fiber.Ctx) error {
	return rh.setVerifyState(c, true)
}

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
