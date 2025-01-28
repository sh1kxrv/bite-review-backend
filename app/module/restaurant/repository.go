package restaurant

import (
	"bitereview/database"
	"bitereview/entity"
	"bitereview/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const RestaurantCollection = "restaurants"

type RestaurantRepository struct {
	repository.CrudRepository[entity.Restaurant]
	Collection *mongo.Collection
}

func NewRestaurantRepository() *RestaurantRepository {
	return &RestaurantRepository{
		Collection:     database.GetCollection(RestaurantCollection),
		CrudRepository: repository.NewCrudRepository[entity.Restaurant](RestaurantCollection),
	}
}

func (rp *RestaurantRepository) UpdateVerifiedStatus(context context.Context, restId primitive.ObjectID, verifiedState bool) error {
	filter := bson.M{"_id": restId}
	update := bson.M{"$set": bson.M{"isVerified": verifiedState}}
	return rp.UpdateBSON(context, filter, update)
}
