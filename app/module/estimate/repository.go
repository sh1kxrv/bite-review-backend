package estimate

import (
	"bitereview/database"
	"bitereview/entity"
	"bitereview/repository"
	"bitereview/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const EstimateCollection = "estimates"

type EstimateRepository struct {
	repository.CrudRepository[entity.Estimate]
	Collection *mongo.Collection
}

func NewEstimateRepository() *EstimateRepository {
	return &EstimateRepository{
		Collection:     database.GetCollection(EstimateCollection),
		CrudRepository: repository.NewCrudRepository[entity.Estimate](EstimateCollection),
	}
}

func (r *EstimateRepository) GetEntitiesByReviewId(
	ctx context.Context, reviewId primitive.ObjectID, limit, offset int64,
) ([]entity.Estimate, error) {
	return utils.CursoredFind[entity.Estimate](
		r.Collection, ctx, bson.M{"reviewId": reviewId}, limit, offset,
	)
}
