package metricsfx

import (
	"bitmagnet-io/bitmagnet/internal/metrics/queuemetrics"
	"bitmagnet-io/bitmagnet/internal/metrics/torrentmetrics"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"queue",
		fx.Provide(
			queuemetrics.New,
			torrentmetrics.New,
		),
	)
}
