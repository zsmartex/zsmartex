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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var _ UserUsecase = (*userUsecase)(nil)

type userUsecase struct {
	commandBus      command.Bus
	usersCollection *mongo.Collection
}

func NewUserUsecase(
	ctx context.Context,
	commandBus command.Bus,
	mongoClient *mongo.Client,
) UserUsecase {
	usersCollection := mongoClient.Database("user").Collection("users")

	usersCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{
				"uid": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	})

	return &userUsecase{
		commandBus:      commandBus,
		usersCollection: usersCollection,
	}
}

func (u *userUsecase) emailExist(ctx context.Context, email string) (found bool, err error) {
	userEmailFilter := bson.M{"email": email}
	userEmailExists := u.usersCollection.FindOne(ctx, userEmailFilter)
	if errors.Is(userEmailExists.Err(), mongo.ErrNoDocuments) {
		return false, nil
	}

	if userEmailExists.Err() != nil {
		return false, errors.Cause(userEmailExists.Err())
	}

	return true, nil
}

func (u *userUsecase) UserExists(ctx context.Context, uid string) (found bool, err error) {
	userUIDFilter := bson.M{"uid": uid}
	userUIDExists := u.usersCollection.FindOne(ctx, userUIDFilter)
	if errors.Is(userUIDExists.Err(), mongo.ErrNoDocuments) {
		return false, nil
	}

	if userUIDExists.Err() != nil {
		return false, errors.Cause(userUIDExists.Err())
	}

	return true, nil
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
