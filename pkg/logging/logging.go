package logging

import (
	"context"

	"github.com/modernice/goes/helper/streams"
	"github.com/zsmartex/pkg/v2/log"
)

// LogErrors logs all errors from the provided channels.
func LogErrors(ctx context.Context, errs ...<-chan error) {
	log.Debugf("Logging errors ...")

	in := streams.FanInContext(ctx, errs...)
	for err := range in {
		log.Error(err)
	}
}
