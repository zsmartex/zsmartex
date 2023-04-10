package users

// import (
// 	"context"

// 	"github.com/google/wire"
// 	eh "github.com/looplab/eventhorizon"
// 	"github.com/looplab/eventhorizon/uuid"
// 	"github.com/zsmartex/pkg/v3/gpa/filters"
// 	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// var UseCaseSet = wire.NewSet(NewUserUseCase)

// type UseCase interface {
// 	Login(ctx context.Context, params LoginParams) (user *userv1.UserORM, err error)
// 	GetUser(ctx context.Context, params GetUserParams) (*userv1.UserORM, error)
// 	CreateUser(context.Context, CreateUserParams) (*userv1.UserORM, error)
// 	UpdateUser(ctx context.Context, params UpdateUserParams) (user *userv1.UserORM, err error)
// }

// type service struct {
// 	client
// }

// func NewUserUseCase(repo eh.ReadRepo) UseCase {
// 	return service{
// 		repo,
// 	}
// }

// type LoginParams struct {
// 	CredentialsType  userv1.UserCredentials_Type
// 	CredentialsValue string
// 	Password         string
// }

// func (s service) Login(ctx context.Context, params LoginParams) (user *userv1.UserORM, err error) {
// 	userCredentials, err := s.userCredentialsRepository.GetUserCredentials(ctx, repo.GetUserCredentialsParams{
// 		Type:  params.CredentialsType,
// 		Value: params.CredentialsValue,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	user, err = s.userRepository.First(
// 		ctx,
// 		filters.WithFieldEqual("user_id", userCredentials.UserId),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(params.Password))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// type GetUserParams struct {
// 	Email string
// 	Phone string
// }

// func (s service) GetUser(ctx context.Context, params GetUserParams) (*userv1.UserORM, error) {
// 	var userCredentialsValue string
// 	var userCredentialsType userv1.UserCredentials_Type

// 	if len(params.Email) > 0 {
// 		userCredentialsValue = params.Email
// 		userCredentialsType = userv1.UserCredentials_EMAIL
// 	} else {
// 		userCredentialsValue = params.Phone
// 		userCredentialsType = userv1.UserCredentials_PHONE
// 	}

// 	mongo

// 	userCredentials, err := s.userCredentialsRepository.GetUserCredentials(ctx, repo.GetUserCredentialsParams{
// 		Value: userCredentialsValue,
// 		Type:  userCredentialsType,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	user, err := s.userRepository.GetUser(ctx, userCredentials.UserId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// type CreateUserParams struct {
// 	Email    string
// 	Phone    string
// 	Password string
// 	Role     string
// 	State    string
// 	Level    int32
// }

// func (s service) CreateUser(ctx context.Context, params CreateUserParams) (*userv1.UserORM, error) {
// 	user := &userv1.UserORM{}

// 	err := s.userRepository.Transaction(func(tx *gorm.DB) (err error) {
// 		user, err := s.userRepository.WithTrx(tx).CreateUser(ctx, repo.CreateUserParams{
// 			Password: params.Password,
// 			Role:     params.Role,
// 			State:    params.State,
// 			Level:    params.Level,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		userCredentialsValue := params.Email
// 		userCredentialsType := userv1.UserCredentials_EMAIL
// 		if len(params.Phone) > 0 {
// 			userCredentialsValue = params.Phone
// 			userCredentialsType = userv1.UserCredentials_PHONE
// 		}

// 		_, err = s.userCredentialsRepository.WithTrx(tx).CreateOrUpdateUserCredentials(ctx, repo.CreateOrUpdateUserCredentialsParams{
// 			UserID: user.Id,
// 			Type:   userCredentialsType,
// 			Value:  userCredentialsValue,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	return user, err
// }

// type UpdateUserParams struct {
// 	UserID         uuid.UUID
// 	Email          string
// 	Phone          string
// 	PasswordDigest string
// 	State          userv1.UserState
// 	Role           userv1.UserRole
// }

// func (s service) UpdateUser(ctx context.Context, params UpdateUserParams) (user *userv1.UserORM, err error) {
// 	err = s.userRepository.Transaction(func(tx *gorm.DB) error {
// 		if len(params.Email) > 0 {
// 			_, err = s.userCredentialsRepository.WithTrx(tx).CreateOrUpdateUserCredentials(ctx, repo.CreateOrUpdateUserCredentialsParams{
// 				UserID: params.UserID,
// 				Type:   userv1.UserCredentials_EMAIL,
// 				Value:  params.Email,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		if len(params.Phone) > 0 {
// 			_, err = s.userCredentialsRepository.WithTrx(tx).CreateOrUpdateUserCredentials(ctx, repo.CreateOrUpdateUserCredentialsParams{
// 				UserID: params.UserID,
// 				Type:   userv1.UserCredentials_PHONE,
// 				Value:  params.Phone,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		user, err = s.userRepository.WithTrx(tx).UpdateUser(ctx, repo.UpdateUserParams{
// 			UserID:   params.UserID,
// 			Password: params.PasswordDigest,
// 			State:    params.State,
// 			Role:     params.Role,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	return user, nil
// }
