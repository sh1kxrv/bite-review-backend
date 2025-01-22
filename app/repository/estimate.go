package repository

import (
	"bitereview/app/database"
	"bitereview/app/schema"

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
