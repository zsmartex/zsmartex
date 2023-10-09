package repo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zsmartex/zsmartex/internal/user/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ UserRepo = (*userRepo)(nil)

type UserRepo interface {
	GetUser(ctx context.Context, queryBy string, queryValue string) (*domain.User, error)
}

type userRepo struct {
	usersCollection *mongo.Collection
}

func NewUserRepo(ctx context.Context, mongoClient *mongo.Client) UserRepo {
	usersCollection := mongoClient.Database("user").Collection("users")

	usersCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{
				"uid": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	})

	return &userRepo{
		usersCollection,
	}
}

func (u *userRepo) GetUser(ctx context.Context, queryBy, queryValue string) (*domain.User, error) {
	res := u.usersCollection.FindOne(ctx, bson.M{queryBy: queryValue})
	if res.Err() != nil {
		return nil, errors.Cause(res.Err())
	}

	var user *domain.User
	err := res.Decode(&user)
	if err != nil {
		return nil, errors.Cause(err)
	}

	return user, nil
}
