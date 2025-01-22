package schema

import (
	"bitereview/app/enum"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID  `bson:"_id" json:"id"`
	Email      string              `bson:"email" json:"email"`
	Username   *string             `bson:"username" json:"username"`
	FirstName  string              `bson:"firstName" json:"firstName"`
	LastName   string              `bson:"lastName" json:"lastName"`
	Password   string              `bson:"password" json:"-"`
	IsVerified bool                `bson:"isVerified" json:"isVerified"`
	VerifiedBy *primitive.ObjectID `bson:"verifiedBy" json:"verifiedBy"`
	LastSeen   time.Time           `bson:"lastSeen" json:"lastSeen"`
	CreatedAt  time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time           `bson:"updatedAt" json:"updatedAt"`
	Role       enum.Role           `bson:"role" json:"role"`
}

func (u *User) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type my User
	return bson.Marshal((*my)(u))
}
