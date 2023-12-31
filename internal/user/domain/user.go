package domain

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"
)

type User struct {
	ID             uuid.UUID `bson:"_id"`
	UID            string    `bson:"uid"`
	Email          string    `bson:"email"`
	PasswordDigest string    `bson:"password_digest"`
	Role           UserRole  `bson:"role"`
	State          UserState `bson:"state"`
	Labels         []Label   `bson:"labels"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at"`
}

func DefaultState() UserState {
	return UserStatePending
}

func DefaultRole() UserRole {
	return UserRoleMember
}

func NewUser(id uuid.UUID, uid, email, passwordDigest string) *User {
	return &User{
		ID:             id,
		UID:            uid,
		Email:          email,
		PasswordDigest: passwordDigest,
		Role:           DefaultRole(),
		State:          DefaultState(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (u User) TableName() string {
	return "users"
}

func (u *User) SetEmail(email string) {
	u.Email = email
	u.UpdatedAt = time.Now()
}

func (u *User) SetPasswordDigest(passwordDigest string) {
	u.PasswordDigest = passwordDigest
	u.UpdatedAt = time.Now()
}

func (u *User) SetRole(role UserRole) {
	u.Role = role
	u.UpdatedAt = time.Now()
}

func (u *User) SetState(state UserState) {
	u.State = state
	u.UpdatedAt = time.Now()
}

func (u *User) AddLabel(label Label) {
	u.Labels = append(u.Labels, label)
	u.UpdatedAt = time.Now()
}

func (u *User) RemoveLabel(label Label) {
	for i, l := range u.Labels {
		if l.Key == label.Key {
			u.Labels = append(u.Labels[:i], u.Labels[i+1:]...)
		}
	}

	u.UpdatedAt = time.Now()
}

func (u *User) GetLabel(key string) (Label, bool) {
	for _, l := range u.Labels {
		if l.Key == key {
			return l, true
		}
	}

	return Label{}, false
}

func (u *User) ProtobufValue() *commonv1.User {
	labels := make([]*commonv1.UserLabel, 0)

	for _, label := range u.Labels {
		labels = append(labels, label.ProtobufValue())
	}

	return &commonv1.User{
		Id:        u.ID.String(),
		Uid:       u.UID,
		Email:     u.Email,
		Role:      u.Role.ProtobufValue(),
		State:     u.State.ProtobufValue(),
		Labels:    labels,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}
