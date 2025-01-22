package handler

import (
	"bitereview/app/enum"
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/middleware"
	"bitereview/app/repository"
	"bitereview/app/schema"
	"bitereview/app/serializer"
	"bitereview/app/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RestaurantHandler struct {
	RestaurantRepo *repository.RestaurantRepository
}

func NewRestaurantHandler(restaurantRepo *repository.RestaurantRepository) *RestaurantHandler {
	return &RestaurantHandler{
		RestaurantRepo: restaurantRepo,
	}
}

func (rh *RestaurantHandler) GetRestaurants(c *fiber.Ctx) error {
	return nil
}

func (rh *RestaurantHandler) GetRestaurantById(c *fiber.Ctx) error {
	restId := c.Params("id")
	parsedId, err := primitive.ObjectIDFromHex(restId)
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	restaurant, err := rh.RestaurantRepo.GetEntityByID(timeoutCtx, parsedId)
	if err != nil {
		return helper.SendError(c, err, errors.MakeRepositoryError("Restaurant"))
	}

	return helper.SendSuccess(c, restaurant)
}

func (rh *RestaurantHandler) CreateRestaurant(c *fiber.Ctx) error {
	data, err := serializer.GetSerializedCreateRestaurant(c)
	if err != nil {
		return err
	}

	_, _, err = serializer.GetJwtUserLocalWithParsedID(c)

	if err != nil {
		return helper.SendError(c, err, errors.Unauthorized)
	}

	restourant := &schema.Restaurant{
		ID:          primitive.NewObjectID(),
		Name:        data.Name,
		Description: data.Description,
		Address:     data.Address,
		Location:    data.Location,
		Country:     data.Country,
		Site:        data.Site,
		IsVerified:  false,
		Metadata:    data.Metadata,
	}

	withTimeout, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	rh.RestaurantRepo.CreateEntity(withTimeout, restourant)

	return nil
}

func (rh *RestaurantHandler) VerifyRestaurant(c *fiber.Ctx) error {
	return nil
}

func (rh *RestaurantHandler) UnverifyRestaurant(c *fiber.Ctx) error {
	return nil
}

func (rh *RestaurantHandler) RegisterRoutes(g fiber.Router) {
	adminRoute := g.Group("/admin/restaurant", middleware.JwtAuthMiddleware, middleware.CreateRoleMiddleware(enum.RoleAdmin))
	adminRoute.Post("/", rh.CreateRestaurant)

	moderRoute := g.Group("/moderator/restaurant", middleware.JwtAuthMiddleware, middleware.CreateRoleMiddleware(enum.RoleModerator))
	moderRoute.Patch("/:id/verify", rh.VerifyRestaurant)
	moderRoute.Patch("/:id/unverify", rh.UnverifyRestaurant)

	// Public
	restRoute := g.Group("/restaurant")

	// TODO: Limit / Offset
	restRoute.Get("/", rh.GetRestaurants)
	restRoute.Get("/:id", rh.GetRestaurantById)
}
