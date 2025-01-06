package telemetryfx

import (
	"bitmagnet-io/bitmagnet/internal/telemetry/httpserver"
	"bitmagnet-io/bitmagnet/internal/telemetry/prometheus"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"telemetry",
		fx.Provide(
			httpserver.New,
			prometheus.New,
		),
	)
}
