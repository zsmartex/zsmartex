package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	UID            string             `bson:"uid"`
	Email          string             `bson:"email"`
	PasswordDigest string             `bson:"password_digest"`
	Role           UserRole           `bson:"role"`
	State          UserState          `bson:"state"`
	Labels         []Label            `bson:"labels"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}
