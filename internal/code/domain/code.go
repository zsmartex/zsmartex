package domain

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Code struct {
	ID             primitive.ObjectID `bson:"_id"`
	UserID         primitive.ObjectID `bson:"user_id"`
	Code           string             `bson:"code"`
	Type           CodeType           `bson:"type"`
	Category       CodeCategory       `bson:"category"`
	EmailIndex     int64              `bson:"email_index,omitempty"`
	EmailEncrypted string             `bson:"email_encrypted,omitempty"`
	PhoneIndex     int64              `bson:"phone_index,omitempty"`
	PhoneEncrypted string             `bson:"phone_encrypted,omitempty"`
	AttemptCount   int64              `bson:"attempt_count"`
	ValidatedAt    time.Time          `bson:"validated_at,omitempty"`
	ExpiredAt      time.Time          `bson:"expired_at"`
	Data           json.RawMessage    `bson:"data,omitempty"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}
