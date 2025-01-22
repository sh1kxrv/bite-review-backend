package repository

import (
	"bitereview/app/database"
	"bitereview/app/schema"

	"go.mongodb.org/mongo-driver/mongo"
)

const RestaurantCollection = "restaurants"

type RestaurantRepository struct {
	CrudRepository[schema.Restaurant]
	Collection *mongo.Collection
}

func NewRestaurantRepository() *RestaurantRepository {
	return &RestaurantRepository{
		Collection:     database.GetCollection(RestaurantCollection),
		CrudRepository: NewCrudRepository[schema.Restaurant](RestaurantCollection),
	}
}
