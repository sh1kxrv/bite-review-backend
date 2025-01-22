package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Score struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	RestaurantID primitive.ObjectID `bson:"restaurantId" json:"restaurantId"`
	UserID       primitive.ObjectID `bson:"userId" json:"userId"`
	Flavor       int                `bson:"flavor" json:"flavor"`
	Presentation int                `bson:"presentation" json:"presentation"`
	Serving      int                `bson:"serving" json:"serving"`
	Temperature  int                `bson:"temperature" json:"temperature"`
	Quality      int                `bson:"quality" json:"quality"`
	Uniqueness   int                `bson:"uniqueness" json:"uniqueness"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
}

func (u *Score) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}

	type my Score
	return bson.Marshal((*my)(u))
}
