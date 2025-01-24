package service

import (
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/param"
	"bitereview/app/repository"
	"bitereview/app/schema"
	"bitereview/app/serializer"
	"bitereview/app/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RestaurantService struct {
	RestaurantRepo *repository.RestaurantRepository
}

func NewRestaurantService(restaurantRepo *repository.RestaurantRepository) *RestaurantService {
	return &RestaurantService{
		RestaurantRepo: restaurantRepo,
	}
}

func (rs *RestaurantService) GetRestaurants(c *fiber.Ctx) error {
	limit, offset := param.GetLimitOffset(c)
	restaurants, err := rs.RestaurantRepo.GetAll(limit, offset)
	if err != nil {
		return helper.SendError(c, err, errors.MakeRepositoryError("Restaurant"))
	}

	return helper.SendSuccess(c, restaurants)
}

func (rs *RestaurantService) GetRestaurantById(c *fiber.Ctx) error {
	restId, err := param.ParamPrimitiveID(c, "id")
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	restaurant, err := rs.RestaurantRepo.GetEntityByID(timeoutCtx, restId)
	if err != nil {
		return helper.SendError(c, err, errors.MakeRepositoryError("Restaurant"))
	}

	return helper.SendSuccess(c, restaurant)
}

func (rs *RestaurantService) CreateRestaurant(c *fiber.Ctx) error {
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
		Description: *data.Description,
		Address:     data.Address,
		City:        data.City,
		Country:     data.Country,
		Site:        data.Site,
		KitchenType: data.KitchenType,
		IsVerified:  false,
		Metadata:    data.Metadata,
	}

	withTimeout, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	rs.RestaurantRepo.CreateEntity(withTimeout, restourant)

	return helper.SendSuccess(c, restourant)
}
