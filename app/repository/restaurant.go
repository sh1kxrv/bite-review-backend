package repository

import (
	"bitereview/app/database"
	"bitereview/app/schema"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (rp *RestaurantRepository) UpdateVerifiedStatus(context context.Context, restId primitive.ObjectID, verifiedState bool) error {
	filter := bson.M{"_id": restId}
	update := bson.M{"$set": bson.M{"isVerified": verifiedState}}
	return rp.UpdateBSON(context, filter, update)
}
