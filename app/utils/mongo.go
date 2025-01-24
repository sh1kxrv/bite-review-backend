package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CursoredFind[T any](collection *mongo.Collection, context context.Context, filter interface{}) ([]T, error) {
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
