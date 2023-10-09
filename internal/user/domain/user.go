package domain

import (
	"time"

	commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (u *User) ProtobufValue() *commonv1.User {
	labels := make([]*commonv1.UserLabel, 0)

	for _, label := range u.Labels {
		labels = append(labels, label.ProtobufValue())
	}

	return &commonv1.User{
		Id:        u.ID.Hex(),
		Uid:       u.UID,
		Email:     u.Email,
		Role:      u.Role.ProtobufValue(),
		State:     u.State.ProtobufValue(),
		Labels:    labels,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}
