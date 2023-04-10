package repo

import (
	"context"

	"github.com/zsmartex/zsmartex/internal/users/domain"
	"github.com/zsmartex/zsmartex/pkg/encryption"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUser(ctx context.Context, params GetUserParams) (*domain.User, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return userRepository{
		collection: collection,
	}
}

type GetUserParams struct {
	Email string
	Phone string
}

func (r userRepository) GetUser(ctx context.Context, params GetUserParams) (*domain.User, error) {
	result := r.collection.FindOne(ctx, bson.M{
		"$or": []map[string]interface{}{
			{
				"credentials.type":  userv1.UserCredentials_EMAIL,
				"credentials.value": encryption.Encrypt(params.Email),
			},
			{
				"credentials.type":  userv1.UserCredentials_PHONE,
				"credentials.value": encryption.Encrypt(params.Phone)},
		},
	})
	if err := result.Err(); err != nil {
		return nil, err
	}

	var userEntity *domain.User
	if err := result.Decode(&userEntity); err != nil {
		return nil, err
	}

	return userEntity, nil
}
