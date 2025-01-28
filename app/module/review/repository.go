package review

import (
	"bitereview/database/mongodb"
	"bitereview/entity"
	"bitereview/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

const ReviewCollection = "restaurants"

type ReviewRepository struct {
	repository.CrudRepository[entity.Review]
	Collection *mongo.Collection
}

func NewReviewRepository(db *mongodb.MongoInstance) *ReviewRepository {
	return &ReviewRepository{
		Collection:     db.GetCollection(ReviewCollection),
		CrudRepository: repository.NewCrudRepository[entity.Review](ReviewCollection, db),
	}
}
