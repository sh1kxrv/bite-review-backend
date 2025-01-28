package restaurant

import (
	"bitereview/entity"
	"bitereview/errors"
	"bitereview/helper"
	"bitereview/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RestaurantService struct {
	restaurantRepo *RestaurantRepository
}

func NewRestaurantService(restaurantRepo *RestaurantRepository) *RestaurantService {
	return &RestaurantService{
		restaurantRepo: restaurantRepo,
	}
}

func (rs *RestaurantService) GetRestaurants(limit, offset int64) (*[]entity.Restaurant, *helper.ServiceError) {
	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	restaurants, err := rs.restaurantRepo.GetAll(timeoutCtx, bson.M{}, limit, offset)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.MakeRepositoryError("Restaurant"))
	}

	return &restaurants, nil
}

func (rs *RestaurantService) GetRestaurantById(id primitive.ObjectID) (*entity.Restaurant, *helper.ServiceError) {
	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	restaurant, err := rs.restaurantRepo.GetEntityByID(timeoutCtx, id)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.MakeRepositoryError("Restaurant"))
	}

	return restaurant, nil
}

func (rs *RestaurantService) CreateRestaurant(data *CreateRestaurantDTO) (*entity.Restaurant, *helper.ServiceError) {
	restourant := &entity.Restaurant{
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

	rs.restaurantRepo.CreateEntity(withTimeout, restourant)

	return restourant, nil
}

func (rs *RestaurantService) UpdateVerifiedStatus(id primitive.ObjectID, verifiedState bool) *helper.ServiceError {
	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	err := rs.restaurantRepo.UpdateVerifiedStatus(timeoutCtx, id, verifiedState)
	if err != nil {
		return helper.NewServiceError(err, errors.MakeRepositoryError("Restaurant"))
	}

	return nil
}
