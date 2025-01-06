package healthfx

import (
	"bitmagnet-io/bitmagnet/internal/health"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"health",
		fx.Provide(
			health.New,
		),
	)
}
