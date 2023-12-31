package repo

import (
	"context"

	"go.uber.org/fx"

	"github.com/zsmartex/zsmartex/internal/user/domain"
	"github.com/zsmartex/zsmartex/pkg/mongo_fx"
	"github.com/zsmartex/zsmartex/pkg/mongo_fx/filters"
)

var ReadModule = fx.Module(
	"repo.ReadRepo",
	fx.Provide(
		NewUserReadRepo,
		fx.Annotate(
			mongo_fx.NewReadRepository[domain.User],
			fx.ParamTags(`name:"users_collection"`),
		),
	),
)

type ReadRepo interface {
	mongo_fx.ReadRepository[domain.User]
	GetUserByUID(ctx context.Context, uid string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type readRepo struct {
	mongo_fx.ReadRepository[domain.User]
}

type userReadRepoParams struct {
	fx.In

	mongo_fx.ReadRepository[domain.User]
}

func NewUserReadRepo(ctx context.Context, params userReadRepoParams) ReadRepo {
	return readRepo{
		ReadRepository: params.ReadRepository,
	}
}

func (r readRepo) GetUserByUID(ctx context.Context, uid string) (*domain.User, error) {
	user, err := r.First(ctx, filters.WithFieldEqual("uid", uid))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r readRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.First(ctx, filters.WithFieldEqual("email", email))
	if err != nil {
		return nil, err
	}

	return user, nil
}
