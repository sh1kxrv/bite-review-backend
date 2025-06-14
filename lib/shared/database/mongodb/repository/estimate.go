package repository

import (
	"context"
	"shared/database/mongodb"
	"shared/database/mongodb/entity"
	"shared/utils"
	"shared/utils/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const EstimateCollection = "estimates"

type EstimateRepository struct {
	repository.CrudRepository[entity.Estimate]
	Collection *mongo.Collection
}

func NewEstimateRepository(db *mongodb.MongoInstance) *EstimateRepository {
	return &EstimateRepository{
		Collection:     db.GetCollection(EstimateCollection),
		CrudRepository: repository.NewCrudRepository[entity.Estimate](EstimateCollection, db),
	}
}

func (r *EstimateRepository) GetEntitiesByReviewId(
	ctx context.Context, reviewId primitive.ObjectID, limit, offset int64,
) ([]entity.Estimate, error) {
	return utils.CursoredFind[entity.Estimate](
		r.Collection, ctx, bson.M{"reviewId": reviewId}, limit, offset,
	)
}
