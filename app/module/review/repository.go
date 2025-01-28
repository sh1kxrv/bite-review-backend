package review

import (
	"bitereview/database"
	"bitereview/entity"
	"bitereview/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

const ReviewCollection = "restaurants"

type ReviewRepository struct {
	repository.CrudRepository[entity.Review]
	Collection *mongo.Collection
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{
		Collection:     database.GetCollection(ReviewCollection),
		CrudRepository: repository.NewCrudRepository[entity.Review](ReviewCollection),
	}
}
