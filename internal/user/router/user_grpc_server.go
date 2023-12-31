package router

import (
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	"github.com/zsmartex/pkg/v2/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zsmartex/zsmartex/internal/user/commands"
	"github.com/zsmartex/zsmartex/internal/user/domain"
	"github.com/zsmartex/zsmartex/internal/user/infras/repo"
	userv1 "github.com/zsmartex/zsmartex/proto/api/user/v1"
	commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"
	servicesv1 "github.com/zsmartex/zsmartex/proto/services/v1"
)

var Module = fx.Module(
	"router.Module",
	fx.Provide(
		NewUserServiceServer,
	),
	fx.Invoke(registerRouterHooks),
)

var _ userv1.UserServiceServer = (*userServiceServer)(nil)
var _ servicesv1.UserServiceServer = (*userServiceServer)(nil)

type UserServiceServer interface {
	userv1.UserServiceServer
	servicesv1.UserServiceServer
}

type userServiceServer struct {
	commandBus *cqrs.CommandBus
	readRepo   repo.ReadRepo
}

func NewUserServiceServer(commandBus *cqrs.CommandBus, readRepo repo.ReadRepo) UserServiceServer {
	return &userServiceServer{
		commandBus: commandBus,
		readRepo:   readRepo,
	}
}

func (s *userServiceServer) GetUser(ctx context.Context, req *servicesv1.GetUserRequest) (*commonv1.User, error) {
	return &commonv1.User{}, nil
}

func (s *userServiceServer) Register(ctx context.Context, req *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	id := uuid.New()
	uid := utils.GenerateUID()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	cmd := commands.NewRegisterCommand(id, uid, req.Email, string(hashedPassword))
	err = s.commandBus.Send(ctx, cmd)
	if err != nil {
		return nil, err
	}

	_, err = s.readRepo.GetUserByUID(ctx, uid)
	if err == nil {
		return &userv1.RegisterResponse{}, status.Error(codes.AlreadyExists, "user already exists")
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	_, err = s.readRepo.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return &userv1.RegisterResponse{}, status.Error(codes.AlreadyExists, "user already exists")
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return &userv1.RegisterResponse{
		Uid:   uid,
		Email: req.Email,
		Role:  string(domain.DefaultRole()),
		State: string(domain.DefaultState()),
	}, nil
}

func (s *userServiceServer) Confirm(context.Context, *userv1.ConfirmRequest) (*userv1.ConfirmResponse, error) {
	return &userv1.ConfirmResponse{}, nil
}

func (s *userServiceServer) Login(context.Context, *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	return &userv1.LoginResponse{}, nil
}

func (s *userServiceServer) GenerateCodeRegister(context.Context, *userv1.GenerateCodeRegisterRequest) (*userv1.GenerateCodeRegisterResponse, error) {
	return &userv1.GenerateCodeRegisterResponse{}, nil
}

func (s *userServiceServer) GenerateCodeLogin(context.Context, *userv1.GenerateCodeLoginRequest) (*userv1.GenerateCodeLoginResponse, error) {
	return &userv1.GenerateCodeLoginResponse{}, nil
}
