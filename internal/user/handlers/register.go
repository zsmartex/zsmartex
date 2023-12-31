package handlers

import (
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/zsmartex/zsmartex/internal/user/commands"
	"github.com/zsmartex/zsmartex/internal/user/events"
	"github.com/zsmartex/zsmartex/internal/user/infras/repo"
	"github.com/zsmartex/zsmartex/pkg/logger"
)

type RegisterHandler struct {
	readRepo repo.ReadRepo
	eventBus *cqrs.EventBus
	logger   *logger.Logger
}

type registerHandlerParams struct {
	fx.In

	ReadRepo repo.ReadRepo
	EventBus *cqrs.EventBus
	Logger   *logger.Logger
}

func NewRegisterHandler(params registerHandlerParams) cqrs.CommandHandler {
	handler := RegisterHandler{
		readRepo: params.ReadRepo,
		eventBus: params.EventBus,
		logger:   params.Logger,
	}

	return cqrs.NewCommandHandler("register", handler.Handle)
}

func (h RegisterHandler) Handle(ctx context.Context, cmd *commands.RegisterCmd) error {
	// Check if user already exists
	_, err := h.readRepo.GetUserByUID(ctx, cmd.UID)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return errors.New("user already exists")
	}

	_, err = h.readRepo.GetUserByEmail(ctx, cmd.Email)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return errors.New("user already exists")
	}

	h.logger.Info(
		"User registered",
		zap.String("uid", cmd.UID),
		zap.String("email", cmd.Email),
	)

	return h.eventBus.Publish(ctx, events.UserCreated{
		ID:             uuid.New(),
		UID:            cmd.UID,
		Email:          cmd.Email,
		PasswordDigest: cmd.PasswordDigest,
	})
}
