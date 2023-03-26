package router

import (
	"context"
	"errors"

	"github.com/google/wire"
	usersUC "github.com/zsmartex/zsmartex/internal/user/usecases/users"
	"github.com/zsmartex/zsmartex/pkg/session"
	userv1 "github.com/zsmartex/zsmartex/proto/api/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ProductGRPCServerSet = wire.NewSet(NewUserGRPCServer)

type userGRPCServer struct {
	userv1.UnimplementedUserServiceServer

	sessionStore *session.Store
	usersUseCase usersUC.UseCase
}

func NewUserGRPCServer(
	grpcServer *grpc.Server,
	sessionStore *session.Store,
	usersUseCase usersUC.UseCase,
) userv1.UserServiceServer {
	svc := userGRPCServer{
		sessionStore: sessionStore,
		usersUseCase: usersUseCase,
	}

	userv1.RegisterUserServiceServer(grpcServer, &svc)

	reflection.Register(grpcServer)

	return &svc
}

func (s *userGRPCServer) Login(ctx context.Context, request *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	user, err := s.usersUseCase.GetUser(ctx, usersUC.GetUserParams{
		Email: request.Email,
		Phone: request.Phone,
	})
	if err != nil {
		return nil, err
	}

	if !user.Authenticate(request.Password) {
		return nil, errors.New("wrong password")
	}

	_, err = s.sessionStore.ApplySession(ctx, user)
	if err != nil {
		return nil, err
	}

	userPb, err := user.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &userv1.LoginResponse{
		User: &userPb,
	}, nil
}
