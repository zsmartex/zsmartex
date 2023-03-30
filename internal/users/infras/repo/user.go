package repo

import (
	"context"

	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/pkg/v3/gpa/filters"
	"github.com/zsmartex/pkg/v3/repository"
	"github.com/zsmartex/pkg/v3/usecase"
	"github.com/zsmartex/pkg/v3/utils"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
	"gorm.io/gorm"
)

type UserRepository interface {
	usecase.IUsecase[userv1.UserORM]

	WithTrx(db *gorm.DB) UserRepository
	GetUser(ctx context.Context, userID uuid.UUID) (*userv1.UserORM, error)
	CreateUser(context.Context, CreateUserParams) (*userv1.UserORM, error)
	UpdateUser(ctx context.Context, params UpdateUserParams) (user *userv1.UserORM, err error)
}

type userRepo struct {
	usecase.Usecase[userv1.UserORM]

	credentialsRepo UserCredentialsRepository
}

func NewUserRepository(db *gorm.DB, credentialsRepo UserCredentialsRepository) UserRepository {
	return userRepo{
		Usecase: usecase.Usecase[userv1.UserORM]{
			Repository: repository.New(db, userv1.UserORM{}),
		},
		credentialsRepo: credentialsRepo,
	}
}

func (r userRepo) WithTrx(tx *gorm.DB) UserRepository {
	r.Repository = r.Repository.WithTrx(tx)
	return r
}

func (r userRepo) GetUser(ctx context.Context, userID uuid.UUID) (*userv1.UserORM, error) {
	user, err := r.First(ctx, filters.WithFieldEqual("id", userID))
	if err != nil {
		return nil, err
	}

	return user, nil
}

type CreateUserParams struct {
	Password string
	Role     string
	State    string
	Level    int32
}

func (r userRepo) CreateUser(ctx context.Context, params CreateUserParams) (*userv1.UserORM, error) {
	user := &userv1.UserORM{
		Uid:            utils.GenerateUID(),
		PasswordDigest: userv1.HashPassword(params.Password),
		Role:           params.Role,
		State:          params.State,
		Level:          params.Level,
	}

	err := r.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, err
}

type UpdateUserParams struct {
	UserID   uuid.UUID
	Password string
	State    userv1.UserState
	Role     userv1.UserRole
}

func (r userRepo) UpdateUser(ctx context.Context, params UpdateUserParams) (user *userv1.UserORM, err error) {
	user, err = r.First(
		ctx,
		filters.WithFieldEqual("user_id", params.UserID),
	)
	if err != nil {
		return nil, err
	}

	userUpdatePayload := &userv1.UserORM{}

	if len(params.Password) > 0 {
		userUpdatePayload.PasswordDigest = userv1.HashPassword(params.Password)
	}

	if params.State > 0 {
		userUpdatePayload.State = userv1.UserState_name[int32(params.State)]
	}

	if params.Role > 0 {
		userUpdatePayload.Role = userv1.UserRole_name[int32(params.Role)]
	}

	err = r.Updates(ctx, user, userUpdatePayload)
	if err != nil {
		return nil, err
	}

	return user, nil
}
