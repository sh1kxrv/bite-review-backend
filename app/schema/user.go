package schema

import (
	"bitereview/app/enum"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID  `bson:"_id" json:"id"`
	ProfileID  *primitive.ObjectID `bson:"profileId" json:"profileId"`
	Email      string              `bson:"email" json:"email"`
	Username   string              `bson:"username" json:"username"`
	Password   string              `bson:"password" json:"-"`
	TelegramID *string             `bson:"telegramId" json:"telegramId"`
	LastSeen   time.Time           `bson:"lastSeen" json:"lastSeen"`
	CreatedAt  time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time           `bson:"updatedAt" json:"updatedAt"`
	Role       enum.Role           `bson:"role" json:"role"`
	IsVerified bool                `bson:"isVerified" json:"isVerified"`
}

func (u *User) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type my User
	return bson.Marshal((*my)(u))
}
