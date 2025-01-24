package repository

import (
	"bitereview/app/database"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CrudRepository[T any] struct {
	Collection *mongo.Collection
}

func NewCrudRepository[T any](name string) CrudRepository[T] {
	return CrudRepository[T]{
		Collection: database.GetCollection(name),
	}
}

func (r *CrudRepository[T]) CreateEntity(ctx context.Context, entity *T) (*T, error) {
	_, err := r.Collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *CrudRepository[T]) GetEntityByID(ctx context.Context, id primitive.ObjectID) (*T, error) {
	var entity T
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *CrudRepository[T]) UpdateBSON(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *CrudRepository[T]) GetAll(ctx context.Context, filter bson.M, limit, offset int64) ([]T, error) {

}
