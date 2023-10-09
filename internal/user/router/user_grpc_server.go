package router

import (
	"context"

	"github.com/zsmartex/zsmartex/internal/user/usecases"
	userv1 "github.com/zsmartex/zsmartex/proto/api/user/v1"
	commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"
	servicesv1 "github.com/zsmartex/zsmartex/proto/services/v1"
)

var _ userv1.UserServiceServer = (*userServiceServer)(nil)
var _ servicesv1.UserServiceServer = (*userServiceServer)(nil)

type UserServiceServer interface {
	userv1.UserServiceServer
	servicesv1.UserServiceServer
}

type userServiceServer struct {
	userUsecase usecases.UserUsecase
}

func NewUserServiceServer(
	userUsecase usecases.UserUsecase,
) UserServiceServer {
	return &userServiceServer{
		userUsecase: userUsecase,
	}
}

func (s *userServiceServer) GetUser(ctx context.Context, req *servicesv1.GetUserRequest) (*commonv1.User, error) {
	user, err := s.userUsecase.GetUser(ctx, req.QueryBy, req.QueryValue)
	if err != nil {
		return nil, err
	}

	return user.ProtobufValue(), err
}

func (s *userServiceServer) Register(ctx context.Context, req *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	user, err := s.userUsecase.RegisterUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &userv1.RegisterResponse{
		Uid:   user.UID,
		Email: req.Email,
		Role:  string(user.Role),
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
