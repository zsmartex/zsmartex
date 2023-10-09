package usecases

import (
	"context"

	"github.com/zsmartex/zsmartex/internal/user/domain"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, email string, password string) (*domain.User, error)
}
