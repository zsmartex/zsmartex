package cqrs_fx

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"
)

var CommandProcessorModule = fx.Module(
	"cqrs_fx.CommandProcessorModule",
	fx.Provide(
		NewCommandProcessor,
	),
	fx.Invoke(registerCommandProcessorHooks),
)

func NewCommandProcessor(subscriber message.Subscriber, router *message.Router, marshaler cqrs.CommandEventMarshaler, logger *Logger) (*cqrs.CommandProcessor, error) {
	return cqrs.NewCommandProcessorWithConfig(
		router,
		cqrs.CommandProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.CommandProcessorGenerateSubscribeTopicParams) (string, error) {
				// we are using queue RabbitMQ config, so we need to have topic per command type
				logger.Info("Subscribing to command", watermill.LogFields{
					"command_name": params.CommandName,
				})

				return topicTransform(params.CommandName), nil
			},
			SubscriberConstructor: func(params cqrs.CommandProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				// we can reuse subscriber, because all commands have separated topics
				return subscriber, nil
			},
			OnHandle: func(params cqrs.CommandProcessorOnHandleParams) error {
				start := time.Now()

				err := params.Handler.Handle(params.Message.Context(), params.Command)

				logger.Info("Command handled", watermill.LogFields{
					"command_name": params.CommandName,
					"duration":     time.Since(start),
					"err":          err,
				})

				return err
			},
			Marshaler: marshaler,
			Logger:    logger,
		},
	)
}

type CommandProcessorHooksParams struct {
	fx.In

	CommandProcessor *cqrs.CommandProcessor
	Handlers         []cqrs.CommandHandler `group:"cqrs_command_handlers"`
}

func registerCommandProcessorHooks(params CommandProcessorHooksParams) error {
	return params.CommandProcessor.AddHandlers(
		params.Handlers...,
	)
}
