package repository

import (
	"bitereview/app/database"
	"bitereview/app/schema"

	"go.mongodb.org/mongo-driver/mongo"
)

const ReviewCollection = "restaurants"

type ReviewRepository struct {
	CrudRepository[schema.Review]
	Collection *mongo.Collection
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{
		Collection:     database.GetCollection(ReviewCollection),
		CrudRepository: NewCrudRepository[schema.Review](ReviewCollection),
	}
}
