package models

import (
	"github.com/charoleizer/thuigsinn/ms-users/pkg/proto/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Status   users.Status       `bson:"status"`
}

func (u *User) TableName() string {
	return "users"
}
