package repo

import (
	"context"
	"errors"

	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/pkg/v3/gpa"
	"github.com/zsmartex/pkg/v3/gpa/filters"
	"github.com/zsmartex/pkg/v3/repository"
	"github.com/zsmartex/pkg/v3/usecase"
	"github.com/zsmartex/pkg/v3/utils"
	"github.com/zsmartex/zsmartex/pkg/encryption"
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

func NewUserCredentialsRepository(db *gorm.DB) UserCredentialsRepository {
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

	fs = append(fs, filters.WithFieldEqual("type", params.Type))
	fs = append(fs, filters.WithFieldEqual("value_index", utils.HashStringCRC32(params.Value)))

	return r.First(ctx, fs...)
}

type CreateOrUpdateUserCredentialsParams struct {
	UserID uuid.UUID
	Type   userv1.UserCredentials_Type
	Value  string
}

func (r userCredentialsRepo) CreateOrUpdateUserCredentials(ctx context.Context, params CreateOrUpdateUserCredentialsParams) (*userv1.UserCredentialsORM, error) {
	userCredentials, err := r.First(ctx, filters.WithFieldEqual("user_id", params.UserID))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userCredentials := &userv1.UserCredentialsORM{
			UserId:         params.UserID,
			ValueIndex:     utils.HashStringCRC32(params.Value),
			ValueEncrypted: encryption.Encrypt(params.Value),
			Type:           userv1.UserCredentials_Type_name[int32(params.Type)],
		}
		err := r.Create(ctx, userCredentials)
		if err != nil {
			return nil, err
		}
	} else if userCredentials != nil {
		userCredentialsUpdate := &userv1.UserCredentialsORM{
			ValueIndex:     utils.HashStringCRC32(params.Value),
			ValueEncrypted: encryption.Encrypt(params.Value),
		}
		userCredentials.ValueEncrypted = encryption.Encrypt(params.Value)
		err := r.Updates(ctx, userCredentials, userCredentialsUpdate)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return userCredentials, nil
}
