package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CursoredFindD[T any](collection *mongo.Collection, context context.Context, filter bson.D) ([]T, error) {
	cursor, err := collection.Find(context, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context)

	var stats []T = make([]T, 0)
	err = cursor.All(context, &stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func CursoredFindM[T any](collection *mongo.Collection, context context.Context, filter bson.M) ([]T, error) {
	cursor, err := collection.Find(context, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context)

	var stats []T = make([]T, 0)
	err = cursor.All(context, &stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
