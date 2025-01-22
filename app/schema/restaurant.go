package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Address   string             `bson:"address" json:"address"`
	Location  string             `bson:"location" json:"location"`
	Country   string             `bson:"country" json:"country"`
	Site      string             `bson:"site" json:"site"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

func (r *Restaurant) MarshalBSON() ([]byte, error) {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = time.Now()
	}
	r.UpdatedAt = time.Now()

	type my Restaurant
	return bson.Marshal((*my)(r))
}
