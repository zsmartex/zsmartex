package domain

import (
	"time"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/mbahjadol/null"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
)

type UserLabel struct {
	Key         string    `json:"key" bson:"key"`
	Value       string    `json:"value" bson:"value"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type UserData struct{}

type UserCredentials struct {
	Type           userv1.UserCredentials_Type `json:"type" bson:"type"`
	ValueIndex     int64                       `json:"value_index" bson:"value_index"`
	ValueEncrypted string                      `json:"value_encrypted" bson:"value_encrypted"`
	CreatedAt      time.Time                   `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at" bson:"updated_at"`
}

type User struct {
	ID             uuid.UUID          `json:"id" bson:"_id"`
	Username       null.String        `json:"username" bson:"username"`
	PasswordDigest string             `json:"password_digest" bson:"password_digest"`
	Level          int32              `json:"level" bson:"level"`
	OTP            bool               `json:"otp" bson:"otp"`
	Role           userv1.UserRole    `json:"role" bson:"role"`
	State          userv1.UserState   `json:"state" bson:"state"`
	Credentials    []*UserCredentials `json:"credentials" bson:"credentials"`
	Labels         []*UserLabel       `json:"labels" bson:"labels"`
	Data           UserData           `json:"data" bson:"data"`
	Version        int                `json:"version"    bson:"version"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
}

var _ = eh.Entity(&User{})
var _ = eh.Versionable(&User{})

// EntityID implements the EntityID method of the eventhorizon.Entity interface.
func (t *User) EntityID() uuid.UUID {
	return t.ID
}

// AggregateVersion implements the AggregateVersion method of the
// eventhorizon.Versionable interface.
func (t *User) AggregateVersion() int {
	return t.Version
}
