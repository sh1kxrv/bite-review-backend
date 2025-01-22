package repository

import (
	"bitereview/app/database"
	"bitereview/app/schema"

	"go.mongodb.org/mongo-driver/mongo"
)

const ScoreCollection = "scores"

type ScoreRepository struct {
	CrudRepository[schema.Score]
	Collection *mongo.Collection
}

func NewScoreRepository() *ScoreRepository {
	return &ScoreRepository{
		Collection:     database.GetCollection(ScoreCollection),
		CrudRepository: NewCrudRepository[schema.Score](ScoreCollection),
	}
}
