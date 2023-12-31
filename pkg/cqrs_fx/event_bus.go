package cqrs_fx

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"
)

var EventBusModule = fx.Module(
	"cqrs_fx.EventBusModule",
	fx.Provide(
		NewEventBus,
	),
)

func NewEventBus(publisher message.Publisher, marshaler cqrs.CommandEventMarshaler, logger *Logger) (*cqrs.EventBus, error) {
	return cqrs.NewEventBusWithConfig(publisher, cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
			return topicTransform(params.EventName), nil
		},
		OnPublish: func(params cqrs.OnEventSendParams) error {
			logger.Info("Publishing event", watermill.LogFields{
				"event_name": params.EventName,
			})

			params.Message.Metadata.Set("published_at", time.Now().String())

			return nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})
}
