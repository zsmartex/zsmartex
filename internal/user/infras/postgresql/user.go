package postgresql

import (
	"context"

	"github.com/zsmartex/pkg/v3/repository"
	"github.com/zsmartex/pkg/v3/usecase"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
	"gorm.io/gorm"
)

type UserRepository interface {
	usecase.IUsecase[userv1.UserORM]

	WithTrx(db *gorm.DB) UserRepository
	CreateUser(context.Context, CreateUserParams) (*userv1.UserORM, error)
}

type userRepo struct {
	usecase.Usecase[userv1.UserORM]

	credentialsRepo UserCredentialsRepository
}

func New(db *gorm.DB, credentialsRepo UserCredentialsRepository) UserRepository {
	return userRepo{
		Usecase: usecase.Usecase[userv1.UserORM]{
			Repository: repository.New(db, userv1.UserORM{}),
		},
	}
}

func (r userRepo) WithTrx(tx *gorm.DB) UserRepository {
	r.Repository = r.Repository.WithTrx(tx)
	return r
}

type CreateUserParams struct {
	Email    string
	Phone    string
	Password string
}

func (r userRepo) CreateUser(ctx context.Context, params CreateUserParams) (*userv1.UserORM, error) {
	user := &userv1.UserORM{}

	err := r.Transaction(func(tx *gorm.DB) (err error) {
		err = r.WithTrx(tx).Create(ctx, user)
		if err != nil {
			return err
		}

		userCredentialsValue := params.Email
		userCredentialsType := userv1.UserCredentials_EMAIL
		if len(params.Phone) > 0 {
			userCredentialsValue = params.Phone
			userCredentialsType = userv1.UserCredentials_PHONE
		}

		_, err = r.credentialsRepo.WithTrx(tx).CreateUserCredentials(ctx, CreateUserCredentialsParams{
			UserID: *user.Id,
			Type:   userCredentialsType,
			Value:  userCredentialsValue,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return user, err
}
