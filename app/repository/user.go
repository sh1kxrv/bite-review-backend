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
	Collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Collection: database.GetCollection(UserCollection),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *schema.User) (*schema.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*schema.User, error) {
	var user schema.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateBSON(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *schema.User) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*schema.User, error) {
	var user schema.User
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateLastSeen(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"lastSeen": time.Now()}})
	if err != nil {
		return err
	}
	return nil
}
