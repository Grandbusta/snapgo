package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID          primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	From        primitive.ObjectID `bson:"from" json:"from"`
	Convo       primitive.ObjectID `bson:"convo" json:"convo"`
	Content     string             `bson:"content" json:"content"`
	ContentType string             `bson:"contentType" json:"contentType"`
	IsRead      string             `bson:"is_read" json:"is_read"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

type Convo struct {
	ID           primitive.ObjectID   `bson:"_id, omitempty" json:"id"`
	Participants []primitive.ObjectID `bson:"participants" json:"participants"`
	CreatedAt    time.Time            `bson:"createdAt" json:"createdAt"`
}
