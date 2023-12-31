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

var PublisherModule = fx.Module(
	"cqrs_fx.PublisherModule",
	fx.Provide(
		NewPublisher,
	),
	fx.Invoke(registerPublisherHook),
)

type publisherParams struct {
	fx.In

	Config config.Nats
	Logger *Logger
}

func NewPublisher(params publisherParams) (message.Publisher, error) {
	return nats.NewPublisher(nats.PublisherConfig{
		URL:               params.Config.URL,
		SubjectCalculator: nats.DefaultSubjectCalculator,
		Marshaler:         nats.JSONMarshaler{},
		JetStream: nats.JetStreamConfig{
			Disabled:       false,
			AutoProvision:  true,
			ConnectOptions: nil,
			SubscribeOptions: []nc.SubOpt{
				nc.DeliverAll(),
				nc.AckExplicit(),
			},
			PublishOptions: nil,
			TrackMsgId:     false,
			AckAsync:       false,
			DurablePrefix:  "",
		},
		NatsOptions: []nc.Option{
			nc.RetryOnFailedConnect(true),
			nc.Timeout(30 * time.Second),
			nc.ReconnectWait(1 * time.Second),
		},
	}, params.Logger)
}

func registerPublisherHook(lc fx.Lifecycle, publisher message.Publisher) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return publisher.Close()
		},
	})
}
