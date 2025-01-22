package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Score struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	ReviewID    primitive.ObjectID `bson:"reviewId" json:"reviewId"`
	Name        string             `bson:"name" json:"name"`
	Value       int                `bson:"value" json:"value"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

func (s *Score) MarshalBSON() ([]byte, error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	type my Score
	return bson.Marshal((*my)(s))
}
