package torznabfx

import (
	"bitmagnet-io/bitmagnet/internal/torznab/adapter"
	"bitmagnet-io/bitmagnet/internal/torznab/httpserver"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"torznab",
		fx.Provide(
			adapter.New,
			httpserver.New,
		),
	)
}
