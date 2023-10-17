package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Username  string             `json:"username"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type PublicUser struct {
	ID       primitive.ObjectID `json:"id"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
}

func (u *User) SaveUser() {

}
