package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Status type= pending, accepted, declined

type Friendship struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Sender    primitive.ObjectID `bson:"sender" json:"sender"`
	Receiver  primitive.ObjectID `bson:"receiver" json:"receiver"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}
