package usecases

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/modernice/goes/command"
	"github.com/zsmartex/pkg/v2/utils"
	"github.com/zsmartex/zsmartex/internal/user/commands"
	"github.com/zsmartex/zsmartex/internal/user/domain"
	"github.com/zsmartex/zsmartex/internal/user/handlers"
	"github.com/zsmartex/zsmartex/internal/user/infras/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var _ UserUsecase = (*userUsecase)(nil)

type userUsecase struct {
	commandBus command.Bus
	userRepo   repo.UserRepo
}

func NewUserUsecase(
	ctx context.Context,
	commandBus command.Bus,
	userRepo repo.UserRepo,
) UserUsecase {
	return &userUsecase{
		commandBus: commandBus,
		userRepo:   userRepo,
	}
}

func (u *userUsecase) emailExist(ctx context.Context, email string) (found bool, err error) {
	_, err = u.userRepo.GetUser(ctx, "email", email)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userUsecase) UserExists(ctx context.Context, uid string) (found bool, err error) {
	_, err = u.userRepo.GetUser(ctx, "uid", uid)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userUsecase) GetUser(ctx context.Context, queryBy, queryValue string) (*domain.User, error) {
	return u.userRepo.GetUser(ctx, queryBy, queryValue)
}

func (u *userUsecase) RegisterUser(ctx context.Context, email string, password string) (*domain.User, error) {
	userUID := utils.GenerateUID()

	foundEmail, err := u.emailExist(ctx, email)
	if err != nil {
		return nil, err
	}

	if foundEmail {
		return nil, errors.New("email already exists")
	}

	foundUid, err := u.UserExists(ctx, userUID)
	if err != nil {
		return nil, errors.Cause(err)
	}

	if foundUid {
		return nil, errors.New("uid already exists")
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Cause(err)
	}

	cmd := command.New(string(commands.RegisterUserCmd), commands.RegisterUserData{
		UID:            userUID,
		Email:          email,
		PasswordDigest: string(passwordBytes),
	}, command.Aggregate(handlers.UserAggregate, uuid.New()))
	if err := u.commandBus.Dispatch(ctx, cmd.Any()); err != nil {
		return nil, errors.Cause(err)
	}

	return &domain.User{
		UID:   userUID,
		Email: email,
		Role:  domain.UserRoleMember,
		State: domain.UserStatePending,
	}, nil
}
