package cqrs_fx

import (
	"strings"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"go.uber.org/fx"
)

var MarshalerModule = fx.Module(
	"cqrs_fx.MarshalerModule",
	fx.Provide(
		NewMarshaler,
	),
)

func NewMarshaler() cqrs.CommandEventMarshaler {
	return cqrs.JSONMarshaler{}
}

func topicTransform(name string) string {
	// make it lower and replace dots with underscores
	return strings.ToLower(strings.ReplaceAll(name, ".", "_"))
}
