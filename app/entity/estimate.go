package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Estimate struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	ReviewID    primitive.ObjectID `bson:"reviewId" json:"reviewId"`
	Name        string             `bson:"name" json:"name"`
	Value       int                `bson:"value" json:"value"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

func (s *Estimate) MarshalBSON() ([]byte, error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	type my Estimate
	return bson.Marshal((*my)(s))
}
