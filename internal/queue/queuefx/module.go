package queuefx

import (
	"bitmagnet-io/bitmagnet/internal/queue/manager"
	"bitmagnet-io/bitmagnet/internal/queue/prometheus"
	"bitmagnet-io/bitmagnet/internal/queue/server"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"queue",
		fx.Provide(
			server.New,
			manager.New,
			prometheus.New,
		),
	)
}
