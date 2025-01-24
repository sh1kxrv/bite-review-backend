package repository

import (
	"bitereview/app/database"
	"bitereview/app/schema"
	"bitereview/app/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const EstimateCollection = "estimates"

type EstimateRepository struct {
	CrudRepository[schema.Estimate]
	Collection *mongo.Collection
}

func NewEstimateRepository() *EstimateRepository {
	return &EstimateRepository{
		Collection:     database.GetCollection(EstimateCollection),
		CrudRepository: NewCrudRepository[schema.Estimate](EstimateCollection),
	}
}

func (r *EstimateRepository) GetEntitiesByReviewId(ctx context.Context, reviewId primitive.ObjectID, limit, offset int) ([]schema.Estimate, error) {
	return utils.CursoredFind[schema.Estimate](
		r.Collection, ctx, bson.M{"reviewId": reviewId},
	)
}
