package repo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/zsmartex/zsmartex/internal/code/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CodeRepo interface {
	GetCode(ctx context.Context, queryBy, queryValue string) (*domain.Code, error)
}

type codeRepo struct {
	codesCollection *mongo.Collection
}

func NewCodeRepo(ctx context.Context, mongoClient *mongo.Client) CodeRepo {
	codesCollection := mongoClient.Database("user").Collection("codes")

	codesCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{})

	return &codeRepo{
		codesCollection,
	}
}

func (r *codeRepo) GetCode(ctx context.Context, queryBy, queryValue string) (*domain.Code, error) {
	return nil, errors.New("not implemented")
}

func (r codeRepo) withPending(userID primitive.ObjectID, codeType domain.CodeType, category domain.CodeCategory) bson.M {
	return bson.M{
		"userId":   userID,
		"type":     codeType,
		"category": category,
		"validated_at": []bson.M{
			{"validated_at": nil},
			{"validated_at": bson.M{"$type": 10}},
		},
		"expired_at": bson.M{
			"$gt": primitive.NewDateTimeFromTime(time.Now()),
		},
	}
}
