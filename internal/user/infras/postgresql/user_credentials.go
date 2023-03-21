package postgresql

import (
	"context"

	"github.com/zsmartex/pkg/v3/repository"
	"github.com/zsmartex/pkg/v3/usecase"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
	"gorm.io/gorm"
)

type UserCredentialsRepository interface {
	WithTrx(tx *gorm.DB) UserCredentialsRepository
	CreateUserCredentials(context.Context, CreateUserCredentialsParams) (*userv1.UserCredentialsORM, error)
}

type userCredentialsRepo struct {
	usecase.Usecase[userv1.UserCredentialsORM]
}

func NewCredentials(db *gorm.DB) UserCredentialsRepository {
	return userCredentialsRepo{
		Usecase: usecase.Usecase[userv1.UserCredentialsORM]{
			Repository: repository.New(db, userv1.UserCredentialsORM{}),
		},
	}
}

func (r userCredentialsRepo) WithTrx(tx *gorm.DB) UserCredentialsRepository {
	r.Repository = r.Repository.WithTrx(tx)
	return r
}

type CreateUserCredentialsParams struct {
	UserID int64
	Type   userv1.UserCredentials_Type
	Value  string
}

func (r userCredentialsRepo) CreateUserCredentials(ctx context.Context, params CreateUserCredentialsParams) (*userv1.UserCredentialsORM, error) {
	userCredentials := &userv1.UserCredentialsORM{
		UserId: &params.UserID,
		Type:   userv1.UserCredentials_Type_name[int32(params.Type)],
	}
	err := r.Create(ctx, userCredentials)

	return userCredentials, err
}
