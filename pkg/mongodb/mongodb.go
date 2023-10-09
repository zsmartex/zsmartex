package mongodb

import (
	"context"
	"fmt"

	"github.com/zsmartex/zsmartex/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, config config.MongoDB) (*mongo.Client, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port))
	clientOptions.SetAuth(options.Credential{
		Username: config.Username,
		Password: config.Password,
	})

	return mongo.Connect(ctx, clientOptions)
}
