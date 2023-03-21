package postgresql

import (
	"context"
	"errors"

	"github.com/zsmartex/pkg/v3/gpa"
	"github.com/zsmartex/pkg/v3/gpa/filters"
	"github.com/zsmartex/pkg/v3/repository"
	"github.com/zsmartex/pkg/v3/usecase"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
	"gorm.io/gorm"
)

type UserCredentialsRepository interface {
	WithTrx(tx *gorm.DB) UserCredentialsRepository
	GetUserCredentials(ctx context.Context, params GetUserCredentialsParams) (*userv1.UserCredentialsORM, error)
	CreateOrUpdateUserCredentials(context.Context, CreateOrUpdateUserCredentialsParams) (*userv1.UserCredentialsORM, error)
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

type GetUserCredentialsParams struct {
	Type  userv1.UserCredentials_Type
	Value string
}

func (r userCredentialsRepo) GetUserCredentials(ctx context.Context, params GetUserCredentialsParams) (*userv1.UserCredentialsORM, error) {
	fs := make([]gpa.Filter, 0)

	if params.Type == userv1.UserCredentials_EMAIL {

	} else {

	}

	return r.First(ctx, fs...)
}

type CreateOrUpdateUserCredentialsParams struct {
	UserID int64
	Type   userv1.UserCredentials_Type
	Value  string
}

func (r userCredentialsRepo) CreateOrUpdateUserCredentials(ctx context.Context, params CreateOrUpdateUserCredentialsParams) (*userv1.UserCredentialsORM, error) {
	userCredentials, err := r.First(ctx, filters.WithFieldEqual("user_id", params.UserID))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// do create
	} else if userCredentials != nil {
		// do update
	} else {
		return nil, err
	}

	return userCredentials, err
}
