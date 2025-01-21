package repository

import (
	"bitereview/app/database"
	"bitereview/app/schema"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserCollection = "users"

type UserRepository struct {
	CrudRepository[schema.User]
	Collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Collection:     database.GetCollection(UserCollection),
		CrudRepository: NewCrudRepository[schema.User](UserCollection),
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*schema.User, error) {
	var user schema.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *schema.User) error {
	return r.UpdateBSON(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
}

func (r *UserRepository) UpdateLastSeen(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"lastSeen": time.Now()}})
	if err != nil {
		return err
	}
	return nil
}
