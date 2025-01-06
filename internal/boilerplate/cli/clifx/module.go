package clifx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/cli"
	"bitmagnet-io/bitmagnet/internal/boilerplate/cli/args"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"cli",
		fx.Provide(args.New),
		fx.Provide(cli.New),
	)
}
