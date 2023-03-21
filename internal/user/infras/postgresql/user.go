package postgresql

import (
	"context"

	"github.com/zsmartex/pkg/v3/gpa/filters"
	"github.com/zsmartex/pkg/v3/repository"
	"github.com/zsmartex/pkg/v3/usecase"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
	"golang.org/x/crypto/bcrypt"
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
		user.PasswordDigest = hashPassword(params.Password)

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

		_, err = r.credentialsRepo.WithTrx(tx).CreateOrUpdateUserCredentials(ctx, CreateOrUpdateUserCredentialsParams{
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

type LoginParams struct {
	CredentialsType  userv1.UserCredentials_Type
	CredentialsValue string
	Password         string
}

func (r userRepo) Authentication(ctx context.Context, params LoginParams) (user *userv1.UserORM, err error) {
	userCredentials, err := r.credentialsRepo.GetUserCredentials(ctx, GetUserCredentialsParams{
		Type:  params.CredentialsType,
		Value: params.CredentialsValue,
	})
	if err != nil {
		return nil, err
	}

	user, err = r.First(
		ctx,
		filters.WithFieldEqual("user_id", userCredentials.UserId),
	)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(params.Password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

type UpdateUserParams struct {
	UserID           int64
	CredentialsType  userv1.UserCredentials_Type
	CredentialsValue string
	Password         string
	State            userv1.UserState
	Role             userv1.UserRole
}

func (r userRepo) UpdateUser(ctx context.Context, params UpdateUserParams) (user *userv1.UserORM, err error) {
	user, err = r.First(
		ctx,
		filters.WithFieldEqual("user_id", params.UserID),
	)
	if err != nil {
		return nil, err
	}

	err = r.Transaction(func(tx *gorm.DB) error {
		if len(params.CredentialsValue) > 0 {
			_, err = r.credentialsRepo.WithTrx(tx).CreateOrUpdateUserCredentials(ctx, CreateOrUpdateUserCredentialsParams{
				UserID: params.UserID,
				Type:   params.CredentialsType,
				Value:  params.CredentialsValue,
			})
			if err != nil {
				return err
			}
		}

		userUpdatePayload := &userv1.UserORM{}

		if len(params.Password) > 0 {
			userUpdatePayload.PasswordDigest = hashPassword(params.Password)
		}

		if params.State > 0 {
			userUpdatePayload.State = userv1.UserState_name[int32(params.State)]
		}

		if params.Role > 0 {
			userUpdatePayload.Role = userv1.UserState_name[int32(params.Role)]
		}

		return r.WithTrx(tx).Updates(ctx, user, userUpdatePayload)
	})

	return user, nil
}
