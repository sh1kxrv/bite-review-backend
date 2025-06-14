package repository

import (
	"context"
	"shared/database/mongodb"
	"shared/database/mongodb/entity"
	"shared/utils/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const RestaurantCollection = "restaurants"

type RestaurantRepository struct {
	repository.CrudRepository[entity.Restaurant]
	Collection *mongo.Collection
}

func NewRestaurantRepository(db *mongodb.MongoInstance) *RestaurantRepository {
	return &RestaurantRepository{
		Collection:     db.GetCollection(RestaurantCollection),
		CrudRepository: repository.NewCrudRepository[entity.Restaurant](RestaurantCollection, db),
	}
}

func (rp *RestaurantRepository) UpdateVerifiedStatus(context context.Context, restId primitive.ObjectID, verifiedState bool) error {
	filter := bson.M{"_id": restId}
	update := bson.M{"$set": bson.M{"isVerified": verifiedState}}
	return rp.UpdateBSON(context, filter, update)
}
