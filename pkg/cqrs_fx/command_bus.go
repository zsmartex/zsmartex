package cqrs_fx

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"
)

var CommandBusModule = fx.Module(
	"cqrs_fx.CommandBusModule",
	fx.Provide(
		NewCommandBus,
	),
)

func NewCommandBus(publisher message.Publisher, marshaler cqrs.CommandEventMarshaler, logger *Logger) (*cqrs.CommandBus, error) {
	return cqrs.NewCommandBusWithConfig(publisher, cqrs.CommandBusConfig{
		GeneratePublishTopic: func(params cqrs.CommandBusGeneratePublishTopicParams) (string, error) {
			// we are using queue RabbitMQ config, so we need to have topic per command type
			return topicTransform(params.CommandName), nil
		},
		OnSend: func(params cqrs.CommandBusOnSendParams) error {
			logger.Info("Sending command", watermill.LogFields{
				"command_name": params.CommandName,
			})

			params.Message.Metadata.Set("sent_at", time.Now().String())

			return nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})
}
