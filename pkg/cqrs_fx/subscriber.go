package cqrs_fx

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
	nc "github.com/nats-io/nats.go"
	"go.uber.org/fx"

	"github.com/zsmartex/zsmartex/pkg/config"
)

var SubscriberModule = fx.Module(
	"cqrs_fx.SubscriberModule",
	fx.Provide(
		NewSubscriber,
	),
	fx.Invoke(registerSubscriberHook),
)

type subscriberParams struct {
	fx.In

	Config config.Nats
}

func NewSubscriber(config config.Nats, logger *Logger) (message.Subscriber, error) {
	return nats.NewSubscriber(nats.SubscriberConfig{
		URL:               config.URL,
		SubjectCalculator: nats.DefaultSubjectCalculator,
		Unmarshaler:       nats.JSONMarshaler{},
		JetStream: nats.JetStreamConfig{
			Disabled:       true,
			AutoProvision:  true,
			ConnectOptions: nil,
			SubscribeOptions: []nc.SubOpt{
				nc.DeliverAll(),
				nc.AckExplicit(),
			},
			PublishOptions: nil,
			TrackMsgId:     false,
			AckAsync:       true,
			DurablePrefix:  "",
		},
		NatsOptions: []nc.Option{
			nc.RetryOnFailedConnect(true),
			nc.Timeout(30 * time.Second),
			nc.ReconnectWait(1 * time.Second),
		},
	}, logger)
}

func registerSubscriberHook(lc fx.Lifecycle, subscriber message.Subscriber) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return subscriber.Close()
		},
	})
}
