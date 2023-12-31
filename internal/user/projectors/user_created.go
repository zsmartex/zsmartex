package projectors

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"go.uber.org/fx"

	"github.com/zsmartex/zsmartex/internal/user/domain"
	"github.com/zsmartex/zsmartex/internal/user/events"
	"github.com/zsmartex/zsmartex/internal/user/infras/repo"
)

type UserCreatedProjector struct {
	writeRepo repo.WriteRepo
}

type userCreatedProjectorParams struct {
	fx.In

	WriteRepo repo.WriteRepo
}

func NewUserCreatedProjector(params userCreatedProjectorParams) cqrs.EventHandler {
	projector := &UserCreatedProjector{
		writeRepo: params.WriteRepo,
	}

	return cqrs.NewEventHandler("UserCreated", projector.Handle)
}

func (p *UserCreatedProjector) Handle(ctx context.Context, event *events.UserCreated) error {
	user := domain.NewUser(event.ID, event.UID, event.Email, event.PasswordDigest)

	return p.writeRepo.CreateUser(ctx, user)
}
