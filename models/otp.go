package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// would be much better if we use redis.
type Otp struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Phone *string            `json:"phone,omitempty" bson:"phone,omitempty"`
	Otp   *string            `json:"otp,omitempty" bson:"otp,omitempty"`
	Exp   time.Time          `json:"exp,omitempty"`
}
