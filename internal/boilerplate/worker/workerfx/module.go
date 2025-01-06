package workerfx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/worker"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"worker",
		fx.Provide(worker.NewRegistry),
	)
}
