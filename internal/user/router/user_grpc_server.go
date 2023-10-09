package router

import (
	"context"

	"github.com/zsmartex/zsmartex/internal/user/usecases"
	userv1 "github.com/zsmartex/zsmartex/proto/api/user/v1"
)

type UserServiceServer struct {
	userv1.UnimplementedUserServiceServer

	userUsecase usecases.UserUsecase
}

func NewUserServiceServer(
	userUsecase usecases.UserUsecase,
) *UserServiceServer {
	return &UserServiceServer{
		userUsecase: userUsecase,
	}
}

func (s *UserServiceServer) Register(ctx context.Context, req *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
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
