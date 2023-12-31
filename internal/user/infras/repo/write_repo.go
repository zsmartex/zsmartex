package repo

import (
	"context"

	"go.uber.org/fx"

	"github.com/zsmartex/zsmartex/internal/user/domain"
	"github.com/zsmartex/zsmartex/pkg/mongo_fx"
)

var WriteModule = fx.Module(
	"repo.WriteRepo",
	fx.Provide(
		NewUserWriteRepo,
		fx.Annotate(
			mongo_fx.NewWriteRepository[domain.User],
			fx.ParamTags(`name:"users_collection"`),
		),
	),
)

type WriteRepo interface {
	Migrate(ctx context.Context) error
	CreateUser(ctx context.Context, user *domain.User) error
}

type writeRepo struct {
	mongo_fx.WriteRepository[domain.User]
}

type userWriteRepoParams struct {
	fx.In

	mongo_fx.WriteRepository[domain.User]
}

func NewUserWriteRepo(ctx context.Context, params userWriteRepoParams) WriteRepo {
	return writeRepo{
		WriteRepository: params.WriteRepository,
	}
}

func (r writeRepo) Migrate(ctx context.Context) error {
	// _, err := r.usersCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
	// 	{
	// 		Keys: bson.M{
	// 			"uid": 1,
	// 		},
	// 		Options: options.Index().SetUnique(true),
	// 	},
	// 	{
	// 		Keys: bson.M{
	// 			"email": 1,
	// 		},
	// 		Options: options.Index().SetUnique(true),
	// 	},
	// 	{
	// 		Keys: bson.M{
	// 			"role": 1,
	// 		},
	// 		Options: options.Index(),
	// 	},
	// 	{
	// 		Keys: bson.M{
	// 			"state": 1,
	// 		},
	// 		Options: options.Index(),
	// 	},
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (r writeRepo) CreateUser(ctx context.Context, user *domain.User) error {
	return r.WriteRepository.Create(ctx, user)
}
