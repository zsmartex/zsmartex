package cqrs_fx

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"
)

var EventProcessorModule = fx.Module(
	"cqrs_fx.EventProcessorModule",
	fx.Provide(
		NewEventProcessor,
	),
	fx.Invoke(registerEventProcessorHooks),
)

func NewEventProcessor(subscriber message.Subscriber, router *message.Router, marshaler cqrs.CommandEventMarshaler, logger *Logger) (*cqrs.EventProcessor, error) {
	return cqrs.NewEventProcessorWithConfig(
		router,
		cqrs.EventProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.EventProcessorGenerateSubscribeTopicParams) (string, error) {
				return topicTransform(params.EventName), nil
			},
			SubscriberConstructor: func(params cqrs.EventProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				return subscriber, nil
			},
			OnHandle: func(params cqrs.EventProcessorOnHandleParams) error {
				start := time.Now()

				err := params.Handler.Handle(params.Message.Context(), params.Event)

				logger.Info("Event handled", watermill.LogFields{
					"event_name": params.EventName,
					"duration":   time.Since(start),
					"err":        err,
				})

				return err
			},

			Marshaler: marshaler,
			Logger:    logger,
		},
	)
}

type EventProcessorHooksParams struct {
	fx.In

	EventProcessor *cqrs.EventProcessor
	Handlers       []cqrs.EventHandler `group:"cqrs_event_handlers"`
}

func registerEventProcessorHooks(params EventProcessorHooksParams) error {
	return params.EventProcessor.AddHandlers(
		params.Handlers...,
	)
}
