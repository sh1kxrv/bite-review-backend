package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	RestaurantID primitive.ObjectID `bson:"restaurantId" json:"restaurantId"`
	UserID       primitive.ObjectID `bson:"userId" json:"userId"`
	Summary      string             `bson:"summary" json:"summary"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
}

func (u *Review) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}

	type my Review
	return bson.Marshal((*my)(u))
}
